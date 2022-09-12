package main

import (
	"fmt"
	"log"
	"os"

	"shodan"
	// "github.com/SSRemex/go_hacker_study/02-http-api/01-shodan-client/shodan/"
	// "github.com/blackhat-go/bhg/ch-3/shodan/shodan"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: shodan searchterm")
	}
	// 在go环境中配置shodan的apikey
	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)
	info, err := s.APIInfo()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Query Credits: %d\nScan Credits: %d\n\n", info.QueryCredits, info.ScanCredits)
	hostSearch, err := s.HostSearch(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	for _, host := range hostSearch.Matches {
		fmt.Printf("%18s%8d\n", host.IPString, host.Port)
	}

}
