package utils

import (
	"net/http"
)

type remote struct {
	client *http.Client
}

type RemoteConfig struct {
}

func NewRemote(config RemoteConfig) *remote {
	client := &http.Client{
		Timeout: 0,
	}
	return &remote{
		client: client,
	}
}
