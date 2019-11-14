package utils

import (
	"net/http"
	"context"
	"time"
)


/*
	Fait des requetes GET sur un url specifique avec les configurations definies dans RequestConfig
*/

func (remote *remote) GET(config RequestConfig) (*http.Response, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	url, _ := config.parse()
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	
	request = request.WithContext(ctx)

	for k, v := range config.Header {
		request.Header.Add(k, v)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := remote.client.Do(request)
	if err != nil {
		return nil, err
	}
	return res, nil
}
