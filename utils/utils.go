package utils

import (
	"io"
	"net/http"
)


/*
	Make resquesr to the judge0 api with the http Method
 */
func MakeRequest(method, contentType, token string, reader io.Reader) (*http.Response, error)  {
	client := new(http.Client)
	request, err := http.NewRequest(method, "https://api.judge0.com/submissions/" + token + "?base64_encoded=false&wait=false", reader)
	request.Header.Set("Content-Type", contentType)
	res, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return res, nil
}
