package web

import (
	"github.com/go-redis/redis/v8"
	"github.com/stormi-li/omi-v1/omi-manager"
)

func NewClient(redisClient *redis.Client, omiWebClient, omiServerClient *manager.Client) *Client {
	return &Client{
		redisClient:   redisClient,
		webManager:    omiWebClient,
		serverManager: omiServerClient,
	}
}

func DisableLog() {
	log_cache = false
}
