package main

import (
	"github.com/go-redis/redis/v8"
	"github.com/stormi-li/omi-v1"
)

var redisAddr = "118.25.196.166:3934"
var password = "12982397StrongPassw0rd"

func main() {
	c := omi.NewMonitor(&redis.Options{Addr: redisAddr, Password: password})
	c.Develop("118.25.196.166:9998")
}
