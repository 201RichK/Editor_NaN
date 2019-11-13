package utils

import (
	"net/http"
	"context"
	"time"
	"errors"
)

func (remote *remote) POST(config RequestConfig) (*http.Response, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	
	url, body := config.parse()
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		panic(err)
	}

	request = request.WithContext(ctx)

	for k, v := range config.Header {
		request.Header.Add(k, v)
	}
	request.Header.Set("Content-Type", "application/json")
	res, err := remote.client.Do(request)
	if err != nil {
		return nil, errors.New("Request Timeout")
	}
	return res, nil
}
