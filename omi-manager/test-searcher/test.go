package main

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/stormi-li/omi-v1"
)

var redisAddr = "118.25.196.166:3934"
var password = "12982397StrongPassw0rd"

func main() {
	searcher := omi.NewConfigManager(&redis.Options{
		Addr:     redisAddr,
		Password: password,
	}).NewSearcher()
	searcher.SearchAndListen("mysql", func(address string, data map[string]string) {
		fmt.Println(address)
	})
}
