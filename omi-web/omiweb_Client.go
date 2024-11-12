package omiweb

import (
	"github.com/go-redis/redis/v8"
	ominager "github.com/stormi-li/omi-v1/omi-manager"
)

type Client struct {
	redisClient   *redis.Client
	webManager    *ominager.Client
	serverManager *ominager.Client
}

func (c *Client) NewWebServer(serverName string, weight int) *WebServer {
	return newWebServer(c.redisClient, c.webManager, c.serverManager, serverName, weight)
}

func (c *Client) GenerateTemplate() {
	copyEmbeddedFiles()
}

func (c *Client) NewProxyServer(serverName string) *ProxyServer {
	return &ProxyServer{
		router:       newRouter(c.webManager.NewSearcher()),
		omiWebClient: c.webManager,
		serverName:   serverName,
	}
}
