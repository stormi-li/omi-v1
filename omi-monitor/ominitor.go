package monitor

import (
	"github.com/go-redis/redis/v8"
	manager "github.com/stormi-li/omi-v1/omi-manager"
)

func NewClient(opts *redis.Options, serverSearcher *manager.Searcher, webSearcher *manager.Searcher, configSearcher *manager.Searcher) *Client {
	return &Client{
		serverSearcher: serverSearcher,
		webSearcher:    webSearcher,
		configSearcher: configSearcher,
		opts:           opts,
	}
}
