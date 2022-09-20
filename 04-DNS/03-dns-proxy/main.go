package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/miekg/dns"
)

func parse(filename string) (map[string]string, error) {
	records := make(map[string]string)
	fh, err := os.Open(filename)
	if err != nil {
		return records, err
	}
	defer fh.Close()
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ",", 2)
		if len(parts) > 2 {
			return records, fmt.Errorf("%s is not valid line", line)
		}
		records[parts[0]] = parts[1]
	}
	log.Println("records set to:")
	for k, v := range records {
		fmt.Printf("%s -> %s\n", k, v)
	}

	return records, scanner.Err()
}

func main() {
	var recordLock sync.RWMutex

	records, err := parse("proxy.config")
	if err != nil {
		panic(err)
	}

	dns.HandleFunc(".", func(w dns.ResponseWriter, req *dns.Msg) {
		fmt.Println("req ==>", req)
		if len(req.Question) == 0 {
			dns.HandleFailed(w, req)
			return
		}

		fqdn := req.Question[0].Name
		fqdn = strings.TrimRight(fqdn, ".")
		parts := strings.Split(fqdn, ".")
		if len(parts) >= 2 {
			fqdn = strings.Join(parts[len(parts)-2:], ".")
		}
		recordLock.RLock()
		match := records[fqdn]
		recordLock.RUnlock()

		fmt.Println("match ==> ", match)

		if match == "" {
			dns.HandleFailed(w, req)
			return
		}

		resp, err := dns.Exchange(req, match)
		fmt.Println(resp)
		if err != nil {
			dns.HandleFailed(w, req)
			return
		}
		if err := w.WriteMsg(resp); err != nil {
			dns.HandleFailed(w, req)
		}

	})

	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGUSR1)

		for sig := range sigs {
			switch sig {
			case syscall.SIGUSR1:
				log.Println("SIGUSER1: reloading records")
				recordLock.Lock()
				parse("proxy.config")
				recordLock.Unlock()
			}
		}

	}()

	log.Fatal(dns.ListenAndServe(":53", "udp", nil))

}
