package monitor

import "github.com/stormi-li/omi-v1/omi-manager"

func NewClient(serverSearcher *manager.Searcher, webSearcher *manager.Searcher, configSearcher *manager.Searcher) *Client {
	return &Client{
		serverSearcher: serverSearcher,
		webSearcher:    webSearcher,
		configSearcher: configSearcher,
	}
}
