package main

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/stormi-li/omi-v1"
)

var redisAddr = "118.25.196.166:3934"
var password = "12982397StrongPassw0rd"

func main() {
	searcher := omi.NewServerManager(&redis.Options{Addr: "localhost:6379"}).NewSearcher()
	searcher.SearchAndListen("hello_server", func(address string, data map[string]string) {
		fmt.Println(address)
	})
}
