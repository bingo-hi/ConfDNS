package dns

import (
	"confdns/internal/logger"
	"bytes"
	"crypto/tls"
	"io"
	"net/http"
	"time"

	"github.com/miekg/dns"
)

func forwardToUpstream(w dns.ResponseWriter, r *dns.Msg, upstream string) {
	c := new(dns.Client)
	c.Timeout = 5 * time.Second

	in, _, err := c.Exchange(r, upstream)
	if err != nil {
		logger.LogFatalf("Upstream error: %v", err)
		dns.HandleFailed(w, r)
		return
	}
	w.WriteMsg(in)
}

func resolveViaDoH(domain string, dohURL string, originalID uint16) (*dns.Msg, error) {
	logger.LogDebugf("Preparing to resolve domain: %s via DoH: %s", domain, dohURL)

	msg := new(dns.Msg)
	msg.SetQuestion(dns.Fqdn(domain), dns.TypeA)
	logger.LogDebugf("Constructed DNS query message: %+v", msg)

	dnsQuery, err := msg.Pack()
	if err != nil {
		logger.LogErrorf("Failed to pack DNS message: %v", err)
		return nil, err
	}
	logger.LogDebugf("DNS query packed successfully, length: %d bytes", len(dnsQuery))

	req, err := http.NewRequest("POST", dohURL, bytes.NewReader(dnsQuery))
	if err != nil {
		logger.LogErrorf("Failed to create HTTP request: %v", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/dns-message")
	logger.LogDebugf("Set header Content-Type: application/dns-message")

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // Skip certificate verification for testing
		},
	}

	logger.LogDebugf("Sending POST request to DoH server: %s", dohURL)
	resp, err := client.Do(req)
	if err != nil {
		logger.LogErrorf("DoH request failed: %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	logger.LogDebugf("Received DoH response, status code: %d", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.LogErrorf("Failed to read DoH response body: %v", err)
		return nil, err
	}
	logger.LogDebugf("Successfully read DoH response body, length: %d bytes", len(body))

	dnsResp := new(dns.Msg)
	if err := dnsResp.Unpack(body); err != nil {
		logger.LogErrorf("Failed to unpack DNS response: %v", err)
		return nil, err
	}

	logger.LogDebugf("DNS response unpacked successfully, original ID: %d, will set to original request ID: %d", dnsResp.Id, originalID)
	dnsResp.Id = originalID

	logger.LogDebugf("DoH resolution succeeded, response content: %+v", dnsResp)
	return dnsResp, nil
}

// resolveViaDoT sends a DNS query over TLS (DoT) to the specified DoT server address (e.g., "1.1.1.1:853").
func resolveViaDoT(domain string, dotAddr string, originalID uint16) (*dns.Msg, error) {
	logger.LogDebugf("Resolving domain via DoT: %s using server: %s", domain, dotAddr)

	// 构造 DNS 查询消息
	msg := new(dns.Msg)
	msg.SetQuestion(dns.Fqdn(domain), dns.TypeA)
	logger.LogDebugf("📦 Constructed DNS query message: %+v", msg)

	/*
		tlsConfig := &tls.Config{
			InsecureSkipVerify: true, // 测试环境使用，正式环境建议验证证书
			ServerName:         "",   // 如果你知道 DoT 服务端的域名可填入，用于证书验证
		}
	*/

	// 建立 TLS 连接
	conn, err := dns.DialTimeout("tcp-tls", dotAddr, 5*time.Second)
	if err != nil {
		logger.LogErrorf("Failed to connect to DoT server %s: %v", dotAddr, err)
		return nil, err
	}
	defer conn.Close()

	// 设置 TLS 配置（强制启用）
	if tlsConn, ok := conn.Conn.(*tls.Conn); ok {
		tlsConn.Handshake()
	}

	// 发送 DNS 查询
	err = conn.WriteMsg(msg)
	if err != nil {
		logger.LogErrorf("Failed to write DNS message to DoT server: %v", err)
		return nil, err
	}

	// 读取响应
	resp, err := conn.ReadMsg()
	if err != nil {
		logger.LogErrorf("Failed to read DNS response from DoT server: %v", err)
		return nil, err
	}

	// 设置响应 ID 为原始请求 ID
	resp.Id = originalID
	logger.LogDebugf("Successfully resolved via DoT: %+v", resp)
	return resp, nil
}
