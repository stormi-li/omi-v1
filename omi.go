package omi

import (
	"github.com/go-redis/redis/v8"
	manager "github.com/stormi-li/omi-v1/omi-manager"
	monitor "github.com/stormi-li/omi-v1/omi-monitor"
	web "github.com/stormi-li/omi-v1/omi-web"
)

func NewServerManager(opts *redis.Options) *manager.Client {
	return manager.NewServerManager(opts)
}

func NewWebManager(opts *redis.Options) *manager.Client {
	return manager.NewWebManager(opts)
}

func NewConfigManager(opts *redis.Options) *manager.Client {
	return manager.NewConfigManager(opts)
}

func NewWebClient(opts *redis.Options) *web.Client {
	return web.NewClient(redis.NewClient(opts), NewWebManager(opts), NewServerManager(opts))
}

func NewMonitor(opts *redis.Options) *monitor.Client {
	return monitor.NewClient(opts,NewServerManager(opts).NewSearcher(), NewWebManager(opts).NewSearcher(), NewConfigManager(opts).NewSearcher())
}
