package manager

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/stormi-li/omipc-v1"
)

type Client struct {
	redisClient *redis.Client
	namespace   string
	serverType  string
	omipcClient *omipc.Client
}

func newClient(redisClient *redis.Client, serverType string, prefix string) *Client {
	return &Client{
		redisClient: redisClient,
		namespace:   prefix,
		serverType:  serverType,
		omipcClient: omipc.NewClient(redisClient.Options()),
	}
}

func NewConfigManager(opts *redis.Options) *Client {
	return newClient(redis.NewClient(opts), Config, Prefix_Config)
}

func NewServerManager(opts *redis.Options) *Client {
	return newClient(redis.NewClient(opts), Server, Prefix_Server)
}

func NewWebManager(opts *redis.Options) *Client {
	return newClient(redis.NewClient(opts), Web, Prefix_Web)
}


func (c *Client) NewRegister(serverName, address string) *Register {
	return &Register{
		redisClient: c.redisClient,
		serverName:  serverName,
		address:     address,
		Data:        map[string]string{},
		namespace:   c.namespace,
		ctx:         context.Background(),
		omipcClient: c.omipcClient,
		key:         c.namespace + serverName + namespace_separator + address,
	}
}

func (c *Client) NewSearcher() *Searcher {
	return &Searcher{
		redisClient: c.redisClient,
		namespace:   c.namespace,
		ctx:         context.Background(),
	}
}
