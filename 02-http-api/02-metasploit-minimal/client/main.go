package main

import (
	"fmt"
	"log"
	"msf/rpc"
)

func main() {
	host := "192.168.199.129:55552"
	pass := "123456"
	user := "msf"

	msf, err := rpc.New(host, user, pass)
	if err != nil {
		log.Panicln(err)
	}

	defer msf.Logout()

	sessions, err := msf.SessionList()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("Sessions:")
	for _, session := range sessions {
		fmt.Printf("%5d %s\n", session.ID, session.Info)
	}

}
