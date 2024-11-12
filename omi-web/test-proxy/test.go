package main

import (
	"github.com/go-redis/redis/v8"
	"github.com/stormi-li/omi-v1"
)

var redisAddr = "118.25.196.166:3934"
var password = "12982397StrongPassw0rd"

func main() {
	omiweb := omi.NewWebClient(&redis.Options{Addr: redisAddr, Password: password})
	ps := omiweb.NewProxyServer("http代理")
	ps.SetCache("./cache", 1024*1024*1024)
	ps.StartHttpProxy("118.25.196.166:80")
}
