package ominitor

import ominager "github.com/stormi-li/omi-v1/omi-manager"

func NewClient(serverSearcher *ominager.Searcher, webSearcher *ominager.Searcher, configSearcher *ominager.Searcher) *Client {
	return &Client{
		serverSearcher: serverSearcher,
		webSearcher:    webSearcher,
		configSearcher: configSearcher,
	}
}
