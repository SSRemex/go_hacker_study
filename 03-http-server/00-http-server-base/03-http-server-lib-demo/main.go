package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type badAuth struct {
	Username string
	Password string
}

func (b *badAuth) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")
	if username != b.Username || password != b.Password {
		http.Error(w, "Unauthorized ", 401)
		return
	}
	// 定义一个请求上下文，用来将变量数据传送给下一个handler func，通过r.Context进行携带
	// 初始化请求
	ctx := context.WithValue(r.Context(), "username", username)
	// 携带
	r = r.WithContext(ctx)
	next(w, r)
}

func hello(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)
	fmt.Fprintf(w, "Hi %s\n", username)

}

func main() {
	// 路由
	r := mux.NewRouter()
	r.HandleFunc("/hello", hello).Methods("GET")
	// 中间件
	n := negroni.Classic()
	n.Use(&badAuth{
		Username: "admin",
		Password: "admin",
	})

	n.UseHandler(r)

	http.ListenAndServe(":8877", n)
}
