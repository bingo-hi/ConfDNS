package dns

import (
	"confdns/internal/config"
	"confdns/internal/logger"
	"net"
	"strings"

	"github.com/miekg/dns"
)

func handleDNSRequest(w dns.ResponseWriter, r *dns.Msg, cfg *config.Config) {
	if len(r.Question) == 0 {
		logger.LogErrorf("Received empty DNS question")
		dns.HandleFailed(w, r)
		return
	}

	q := r.Question[0]
	clientIP, _, _ := net.SplitHostPort(w.RemoteAddr().String())
	logger.LogDebugf("💡 From %s DNS request: %s [%s]", clientIP, q.Name, dns.TypeToString[q.Qtype])

	domain := strings.ToLower(q.Name)

	// 1. Block list check
	for _, blocked := range cfg.BlockList {
		if strings.Contains(domain, blocked) {
			logger.LogInfof("Blocked domain: %s", domain)
			resp := new(dns.Msg)
			resp.SetReply(r)
			resp.Rcode = dns.RcodeNameError
			_ = w.WriteMsg(resp)
			return
		}
	}

	// 2. DoH Full or Partial Mode
	if cfg.UseDoH || domainInList(domain, cfg.DoHRequests) {
		logger.LogDebugf("Resolving %s via DoH (%s)", domain, cfg.DoHURL)
		resp, err := resolveViaDoH(domain, cfg.DoHURL, r.Id)
		if err != nil {
			logger.LogErrorf("DoH resolve failed: %v", err)
			dns.HandleFailed(w, r)
			return
		}
		_ = w.WriteMsg(resp)
		return
	}

	// 3. DoT Full or Partial Mode
	if cfg.UseDoT || domainInList(domain, cfg.DoTRequests) {
		logger.LogDebugf("Resolving %s via DoT (%s)", domain, cfg.DoTURL)
		resp, err := resolveViaDoT(domain, cfg.DoTURL, r.Id)
		if err != nil {
			logger.LogErrorf("DoT resolve failed: %v", err)
			dns.HandleFailed(w, r)
			return
		}
		_ = w.WriteMsg(resp)
		return
	}

	// 4. Default fallback: UDP/TCP Upstream
	logger.LogDebugf("Fallback to traditional upstream for domain: %s", domain)
	forwardToUpstream(w, r, cfg.UpstreamDNS)
}

func domainInList(domain string, domainList []string) bool {
	for _, d := range domainList {
		if strings.Contains(domain, strings.ToLower(d)) {
			return true
		}
	}
	return false
}
