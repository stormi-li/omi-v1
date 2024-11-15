package monitor

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-redis/redis/v8"
	manager "github.com/stormi-li/omi-v1/omi-manager"
)

type NodeManager struct {
	serverSearcher *manager.Searcher
	webSearcher    *manager.Searcher
	configSearcher *manager.Searcher
	opts           *redis.Options
}

func NewManager(opts *redis.Options, serverSearcher *manager.Searcher, webSearcher *manager.Searcher, configSearcher *manager.Searcher) *NodeManager {
	return &NodeManager{
		serverSearcher: serverSearcher,
		webSearcher:    webSearcher,
		configSearcher: configSearcher,
		opts:           opts,
	}
}

func (manager *NodeManager) GetServerNodes() map[string]map[string]map[string]string {
	return manager.serverSearcher.SearchAllServers()
}

func (manager *NodeManager) GetWebNodes() map[string]map[string]map[string]string {
	return manager.webSearcher.SearchAllServers()
}

func (manager *NodeManager) GetConfigNodes() map[string]map[string]map[string]string {
	return manager.configSearcher.SearchAllServers()
}

func (manager *NodeManager) Handler(w http.ResponseWriter, r *http.Request) {
	// 获取请求的路径并去掉开头的 '/'
	path := strings.TrimPrefix(r.URL.Path, "/")
	// 以 '/' 分割路径
	parts := strings.Split(path, "/")

	if parts[0] == command_GetWebNodes {
		w.Write([]byte(toJsonStr(manager.GetWebNodes())))
	}
	if parts[0] == command_GetServerNodes {
		w.Write([]byte(toJsonStr(manager.GetServerNodes())))
	}
	if parts[0] == command_GetConfigNodes {
		w.Write([]byte(toJsonStr(manager.GetConfigNodes())))
	}
	if parts[0] == command_UpdateWeight {
		serverType := r.URL.Query().Get("type")
		name := r.URL.Query().Get("name")
		address := r.URL.Query().Get("address")
		weight, _ := strconv.Atoi(r.URL.Query().Get("weight"))
		manager.updateWeight(serverType, name, address, weight)
	}
}

func (nodeManager *NodeManager) updateWeight(serverType, name, address string, weight int) {
	var register manager.Register
	if serverType == manager.Config {
		register = *manager.NewConfigManager(nodeManager.opts).NewRegister(name, address)
	}
	if serverType == manager.Web {
		register = *manager.NewWebManager(nodeManager.opts).NewRegister(name, address)
	}
	if serverType == manager.Server {
		register = *manager.NewServerManager(nodeManager.opts).NewRegister(name, address)
	}
	register.UpdateWeight(weight)
}

func toJsonStr(nodes map[string]map[string]map[string]string) string {
	res := [][]string{}
	for name, addresses := range nodes {
		for address, details := range addresses {
			weight := details["weight"]
			res = append(res, []string{name, address, weight})
		}
	}
	return sliceToJsonStr(res)
}

func sliceToJsonStr(data [][]string) string {
	jsonStr, _ := json.MarshalIndent(data, " ", "  ")
	return string(jsonStr)
}
