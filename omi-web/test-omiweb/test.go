package main

import (
	"embed"

	"github.com/go-redis/redis/v8"
	"github.com/stormi-li/omi-v1"
)

var redisAddr = "118.25.196.166:3934"
var password = "12982397StrongPassw0rd"

//go:embed static/*
var embedSource embed.FS

func main() {
	omiweb := omi.NewWebClient(&redis.Options{Addr: redisAddr, Password: password})
	ws := omiweb.NewWebServer("118.25.196.166", 1)
	ws.Listen("118.25.196.166:7073")
}