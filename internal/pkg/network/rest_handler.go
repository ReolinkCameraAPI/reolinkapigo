package network

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
)

var (
	headers = map[string][]string{
		"Content-Class": {"application/json"},
		"Accept":        {"application/json"},
	}
)

type RestHandlerProxy struct {
	Type string
	Host string
	Port int
}

type RestHandler struct {
	Host     string
	Port     int
	Endpoint string
	Proxy    *RestHandlerProxy
	HTTPS    bool
}

// Create a new RestHandler object with optional argument using Variadic options pattern for customisation
// Refer to the RestHandlerOption<option_name> functions
// https://stackoverflow.com/a/26326418
func NewRestHandler(host string, options ...func(handler *RestHandler) error) (*RestHandler,
	error) {
	restHandler := &RestHandler{
		Host:     host, // the IP only, default scheme is http
		Port:     0, // no port necessary on default config
		Endpoint: "/cgi-bin/api.cgi", // Default endpoint for the Reolink Camera's
	}

	for _, op := range options {
		err := op(restHandler)
		if err != nil {
			return nil, err
		}
	}
	return restHandler, nil
}

// Change the default port to a custom port.
// Default is unset due to http being the default protocol
func RestHandlerOptionPort(port int) func(rh *RestHandler) error {
	return func(rh *RestHandler) error {
		rh.Port = port
		return nil
	}
}

// Change the default endpoint to a custom endpoint
// Default is "/cgi-bin/api.cgi"
// If for some reason the camera you are using is different, one can update it here.
func RestHandlerOptionEndpoint(endpoint string) func(rh *RestHandler) error {
	return func(rh *RestHandler) error {
		rh.Endpoint = endpoint
		return nil
	}
}

// Add a proxy layer on top of the current connection
func RestHandlerOptionProxy(host string, port int) func(rh *RestHandler) error {
	return func(rh *RestHandler) error {
		rh.Proxy = &RestHandlerProxy{
			Host: host,
			Port: port,
		}
		return nil
	}
}

// Change the default scheme from HTTP to HTTPS
func RestHandlerOptionHttp(https bool) func(rh *RestHandler) error {
	return func(rh *RestHandler) error {
		rh.HTTPS = https
		return nil
	}
}

// Do the http request
// endpoint: the trailing part of the URL after the port.
// method: GET or POST
// data: the byte data package needed to be sent.
func (rh *RestHandler) Request(method string, data []byte) (*http.Response, error) {

	var urlConcat string
	if rh.Port > 0 {
		urlConcat = fmt.Sprintf("%s:%d/%s", rh.Host, rh.Port, rh.Endpoint)
	} else {
		urlConcat = fmt.Sprintf("%s/%s", rh.Host, rh.Endpoint)
	}

	if rh.HTTPS {
		urlConcat = fmt.Sprintf("https://%s", urlConcat)
	} else {
		urlConcat = fmt.Sprintf("http://%s", urlConcat)
	}

	url, err := url.Parse(urlConcat)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url.String(), bytes.NewBuffer(data))

	if err != nil {
		return nil, err
	}

	req.Header = headers

	var client *http.Client

	if rh.Proxy != nil {
		tr := http.DefaultTransport.(*http.Transport).Clone()
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		proxyConcat := fmt.Sprintf("%s://%s:%d", rh.Proxy.Type, rh.Proxy.Host, rh.Port)
		proxyUrl, err := url.Parse(proxyConcat)

		if err != nil {
			return nil, err
		}

		tr.Proxy = http.ProxyURL(proxyUrl)
		client = &http.Client{Transport: tr}
	} else {
		client = &http.Client{}
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
