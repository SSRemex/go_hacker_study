package main

import (
	"log"
	"net"

	"github.com/miekg/dns"
)

func main() {
	dns.HandleFunc(".", func(w dns.ResponseWriter, req *dns.Msg) {
		var resp dns.Msg
		resp.SetReply(req)
		for _, q := range req.Question {
			a := dns.A{
				Hdr: dns.RR_Header{
					Name:   q.Name,
					Rrtype: dns.TypeA,
					Class:  dns.ClassINET,
					Ttl:    0,
				},
				// A解析到的IP
				A: net.ParseIP("192.168.199.138").To4(),
			}
			resp.Answer = append(resp.Answer, &a)
		}
		w.WriteMsg(&resp)
	})
	log.Fatal(dns.ListenAndServe(":53", "udp", nil))

}
