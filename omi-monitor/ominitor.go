package ominitor

import omiclient "github.com/stormi-li/omi-v1/omi-manager"

func NewClient(serverSearcher *omiclient.Searcher, webSearcher *omiclient.Searcher, configSearcher *omiclient.Searcher) *Client {
	return &Client{
		serverSearcher: serverSearcher,
		webSearcher:    webSearcher,
		configSearcher: configSearcher,
	}
}
