package main

import (
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/stormi-li/omi-v1"
)

var redisAddr = "118.25.196.166:3934"
var password = "12982397StrongPassw0rd"

func main() {
	Server()
}

func Monitor() {
	c := omi.NewMonitor(&redis.Options{Addr: redisAddr, Password: password})
	c.Listen("118.25.196.166:9998")
}

func Proxy() {
	omiweb := omi.NewWebClient(&redis.Options{Addr: redisAddr, Password: password})
	ps := omiweb.NewProxyServer("http代理")
	ps.StartHttpProxy("118.25.196.166:80")
}

func Server() {
	serverManager := omi.NewServerManager(&redis.Options{Addr: redisAddr, Password: password})
	register := serverManager.NewRegister("hello_server", 1)

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello", r.URL.Query().Get("name"), ", welcome to use omi")
	})

	register.RegisterAndListen("118.25.196.166:8081", func(port string) {
		http.ListenAndServe(port, nil)
	})
}

func Web() {
	web := omi.NewWebClient(&redis.Options{Addr: redisAddr, Password: password})
	web.GenerateTemplate()
	webServer := web.NewWebServer("118.25.196.166", 1)
	webServer.Listen("118.25.196.166:8848")
}
