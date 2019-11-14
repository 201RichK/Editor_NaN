package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/url"
)


/*
	Configurer la requete
	example:
	url: "http://localhost:8080"
	query: map[string][]string{
		"ID": []string{"0"},
		"Name": []string{"Allou"},
	}
	Header: map[string]string{
		"Content-Type": "application/json",
	}
	Body: User{
		Name: "Allou",
	}
*/

type RequestConfig struct {
	URL    string
	Query  map[string][]string
	Header map[string]string
	Body   interface{}
}

func (config *RequestConfig) parse() (string, io.Reader) {
	b, err := json.Marshal(config.Body)
	if err != nil {
		panic(err)
	}
	bufReader := bytes.NewReader(b)
	u, err := url.Parse(config.URL)
	if err != nil {
		panic(err)
	}
	q := u.Query()
	for k, v := range config.Query {
		for _, t := range v {
			q.Add(k, t)
		}
	}
	u.RawQuery = q.Encode()
	return u.String(), bufReader
}
