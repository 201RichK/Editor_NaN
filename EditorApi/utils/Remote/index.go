package utils

import (
	"net/http"
)

type remote struct {
	client *http.Client
}



/*
	Pour Configurer les options du client
*/
type RemoteConfig struct {
}


/*
	Creer un Client qui fait sur un url specifique
*/
func NewRemote(config RemoteConfig) *remote {
	client := &http.Client{
		Transport: &http.Transport{
			TLSHandshakeTimeout: 0,
		},
	}
	return &remote{
		client: client,
	}
}
