package main

import (
	"github.com/miekg/dns"
)

func main() {
	var msg dns.Msg
	// FQDN 完全限定域名
	fqdn := dns.Fqdn("stacktitan.com")
	msg.SetQuestion(fqdn, dns.TypeA)
	dns.Exchange(&msg, "8.8.8.8:53")

}
