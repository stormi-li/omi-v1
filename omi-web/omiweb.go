package omiweb

import (
	"github.com/go-redis/redis/v8"
	ominager "github.com/stormi-li/omi-v1/omi-manager"
)

func NewClient(redisClient *redis.Client, omiWebClient, omiServerClient *ominager.Client) *Client {
	return &Client{
		redisClient:   redisClient,
		webManager:    omiWebClient,
		serverManager: omiServerClient,
	}
}

func DisableLog() {
	log_cache = false
}
