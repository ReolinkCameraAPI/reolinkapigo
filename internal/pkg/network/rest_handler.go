package network

import (
	"bytes"
	"net/http"
)

var (
	headers = map[string][]string{
		"Content-Class": {"application/json"},
		"Accept":        {"application/json"},
	}
)

type RestHandler struct {
	Host  string
	Port  int
	Proxy string
}

func NewRestHandler(host string, port int) *RestHandler {
	return &RestHandler{
		Host:  host,
		Port:  port,
		Proxy: "",
	}
}

func request(url string, method string, data []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))

	if err != nil {
		return nil, err
	}

	req.Header = headers

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
