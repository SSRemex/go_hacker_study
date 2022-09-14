package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var (
	// 赋值upgrader为可信任来源的websocket通信
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	listenAddr string
	wsAddr     string
	jsTemplate *template.Template
)

func init() {
	flag.StringVar(&listenAddr, "listen-addr", "", "Address to listen on")
	flag.StringVar(&wsAddr, "ws-addr", "", "Address for WebSocket connection")
	flag.Parse()

	var err error
	// 解析js文件给jsTemplate变量
	jsTemplate, err = template.ParseFiles("logger.js")
	if err != nil {
		panic(err)
	}

}

func serveWS(w http.ResponseWriter, r *http.Request) {
	// 将所有请求都升级为websocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "", 500)
		return
	}
	defer conn.Close()
	fmt.Printf("Connection from %s\n", conn.RemoteAddr().String())
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		fmt.Printf("From %s:%s\n", conn.RemoteAddr().String(), string(msg))
	}
}

func serveFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	// 将wsAddr填充至logger.js的模板中
	jsTemplate.Execute(w, wsAddr)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ws", serveWS)
	r.HandleFunc("/k.js", serveFile)
	log.Fatal(http.ListenAndServe(":8877", r))
}
