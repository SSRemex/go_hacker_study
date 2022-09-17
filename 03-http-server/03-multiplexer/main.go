package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
)

var (
	hostProxy = make(map[string]string)
	proxies   = make(map[string]*httputil.ReverseProxy)
)

func init() {
	// 代理映射
	hostProxy["attacker1.com"] = "http://192.168.199.129:10080"
	hostProxy["attacker2.com"] = "http://192.168.199.129:20080"

	for k, v := range hostProxy {
		remote, err := url.Parse(v)
		if err != nil {
			log.Fatal("Unable to parse proxy target")
		}
		// fanxiangdaili
		proxies[k] = httputil.NewSingleHostReverseProxy(remote)
	}
}

func main() {
	r := mux.NewRouter()
	for host, proxy := range proxies {
		r.Host(host).Handler(proxy)
	}
	log.Fatal(http.ListenAndServe(":80", r))
}
