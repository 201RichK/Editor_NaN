package utils

import (
	"net/http"
	"context"
	"time"
)

func (remote *remote) PUT(config RequestConfig) (*http.Response, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	url, body := config.parse()
	request, err := http.NewRequest("PUT", url, body)
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
		return nil, err
	}
	return res, nil
}
