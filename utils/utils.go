package utils

import (
	//"crypto/tls"
	"io"
	"net/http"
	"time"
)

/*
	Make resquest to the judge0 api with the http Method
*/
func MakeRequest(method, contentType, token string, reader io.Reader) (*http.Response, error) {
	client := new(http.Client)
	//dire aux client d'attendre la reponse quelque soit le temp
	client.Transport = &http.Transport{
		//		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // disable verify
		TLSHandshakeTimeout: 0,
	}

	//make request to the juqge API
	request, err := http.NewRequest(method, "https://api.judge0.com/submissions/"+token+"?base64_encoded=false&wait=false", reader)
	time.Sleep(5 * time.Second)
	request.Header.Set("Content-Type", contentType)
	res, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return res, nil
}
