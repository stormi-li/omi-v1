package main

import (
	"github.com/go-redis/redis/v8"
	"github.com/stormi-li/omi-v1"
)

var redisAddr = "118.25.196.166:3934"
var password = "12982397StrongPassw0rd"

func main() {
	omiC := omi.NewConfigManager(&redis.Options{
		Addr:     redisAddr,
		Password: password,
	})
	r := omiC.NewRegister("redis", "118.25.196.166:6379")
	r.UpdateWeight(0)
	r = omiC.NewRegister("redis", "118.25.196.166:6378")
	r.Register(1)
	select {}
}
