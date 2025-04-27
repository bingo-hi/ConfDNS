package dns

import (
	"confdns/internal/config"
	"confdns/internal/logger"

	"github.com/miekg/dns"
)

func StartServer(cfg *config.DNSConfig) {
	dns.HandleFunc(".", func(w dns.ResponseWriter, r *dns.Msg) {
		handleDNSRequest(w, r, cfg)
	})

	server := &dns.Server{Addr: cfg.Listen, Net: "udp"}
	logger.LogInfof("Starting DNS server on %s (upstreamDNS:%s)", cfg.Listen, cfg.UpstreamDNS)

	err := server.ListenAndServe()
	if err != nil {
		logger.LogFatalf("Failed to start DNS server: %v", err)
	}
}
