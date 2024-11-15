package manager

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/stormi-li/omipc-v1"
)

type Register struct {
	redisClient *redis.Client
	serverName  string
	address     string
	weight      int
	Data        map[string]string
	namespace   string
	omipcClient *omipc.Client
	key         string
	ctx         context.Context
}

type Options struct {
	Addr       string
	ServerName string
}

func (register *Register) Register(weight int) {
	go register.registerHandler(weight)
	go register.commandHandler()
	log.Println("register server for", register.serverName+"["+register.address+"]", "is starting")
}

func (register *Register) registerHandler(weight int) {
	register.weight = weight
	for {
		register.Data["weight"] = strconv.Itoa(register.weight)
		jsonStrData := mapToJsonStr(register.Data)
		register.redisClient.Set(register.ctx, register.key, jsonStrData, config_expire_time)
		time.Sleep(config_expire_time / 2)
	}
}

func (register *Register) commandHandler() {
	channel := register.namespace + register.serverName + namespace_separator + register.address
	register.omipcClient.Listen(channel, func(message string) bool {
		parts := strings.Split(message, namespace_separator)
		if len(parts) > 1 && parts[0] == command_update_weight {
			weight, _ := strconv.Atoi(parts[1])
			register.weight = weight
		}
		return true
	})
}

func (register *Register) UpdateWeight(weight int) {
	val := command_update_weight + namespace_separator + strconv.Itoa(weight)
	register.omipcClient.Notify(register.key, val)
}

func (register *Register) RegisterAndListen(weight int, handler func(port string)) {
	register.Register(weight)
	handler(":" + strings.Split(register.address, ":")[1])
}
