package omi

import (
	"github.com/go-redis/redis/v8"
	"github.com/stormi-li/omi-v1/omi-manager"
	"github.com/stormi-li/omi-v1/omi-monitor"
	"github.com/stormi-li/omi-v1/omi-web"
)

func NewServerManager(opts *redis.Options) *manager.Client {
	return manager.NewClient(redis.NewClient(opts), manager.Server, manager.Prefix_Server)
}

func NewWebManager(opts *redis.Options) *manager.Client {
	return manager.NewClient(redis.NewClient(opts), manager.Web, manager.Prefix_Web)
}

func NewConfigManager(opts *redis.Options) *manager.Client {
	return manager.NewClient(redis.NewClient(opts), manager.Config, manager.Prefix_Config)
}

func NewWebClient(opts *redis.Options) *web.Client {
	return web.NewClient(redis.NewClient(opts), NewWebManager(opts), NewServerManager(opts))
}

func NewMonitor(opts *redis.Options) *monitor.Client {
	return monitor.NewClient(NewServerManager(opts).NewSearcher(), NewWebManager(opts).NewSearcher(), NewConfigManager(opts).NewSearcher())
}
