package network

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"golang.org/x/net/proxy"
	"io/ioutil"
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
	Type     Proxy
	Protocol *Protocol
	Host     string
	Port     int
	Auth     *RestHandlerProxyAuth
}

type RestHandlerProxyAuth struct {
	Username string
	Password string
}

type RestHandler struct {
	Host     string
	Port     int
	Endpoint string
	Proxy    *RestHandlerProxy
	HTTPS    bool
	*RestAuth
}

type RestAuth struct {
	Username string
	Password string
	Token    string
}

type OptionRestHandler func(*RestHandler)

// Create a new RestHandler object with optional argument using Variadic options pattern for customisation
// Refer to the RestHandlerOption<option_name> functions
// RestHandler is used to wrap the http package and give a cleaner more defined scope which the person
// implementing the library will have full control over.
// https://stackoverflow.com/a/26326418
func NewRestHandler(host string, options ...OptionRestHandler) *RestHandler {
	restHandler := &RestHandler{
		Host:     host,               // the IP only, default scheme is http
		Port:     0,                  // no port necessary on default config
		Endpoint: "cgi-bin/api.cgi", // Default endpoint for the Reolink Camera's
		Proxy: nil,
		HTTPS: false,
		RestAuth: &RestAuth{
			Username: "",
			Password: "",
			Token:    "",
		},
	}

	for _, op := range options {
		op(restHandler)
	}

	return restHandler
}

// Change the default port to a custom port.
// Default is unset due to http being the default protocol
func RestHandlerOptionPort(port int) OptionRestHandler {
	return func(rh *RestHandler) {
		rh.Port = port
	}
}

// Change the default endpoint to a custom endpoint
// Default is "/cgi-bin/api.cgi"
// If for some reason the camera you are using is different, one can update it here.
func RestHandlerOptionEndpoint(endpoint string) OptionRestHandler {
	return func(rh *RestHandler) {
		rh.Endpoint = endpoint
	}
}

// Add a proxy layer on top of the current connection
func RestHandlerOptionProxy(proxy Proxy, host string, port int, auth *RestHandlerProxyAuth,
	protocol *Protocol) OptionRestHandler {
	return func(rh *RestHandler) {
		rh.Proxy = &RestHandlerProxy{
			Host:     host,
			Port:     port,
			Type:     proxy,
			Auth:     auth,
			Protocol: protocol,
		}
	}
}

// Change the default scheme from HTTP to HTTPS
func RestHandlerOptionHttp(https bool) OptionRestHandler {
	return func(rh *RestHandler) {
		rh.HTTPS = https
	}
}

func (rh *RestHandler) SetUsernamePassword(username string, password string) {
	rh.Username = username
	rh.Password = password
}

func (rh *RestHandler) SetToken(token string) {
	rh.Token = token
}

// Do the http request
// endpoint: the trailing part of the URL after the port.
// method: GET or POST
// payload: the json data
// auth: alters the request to include auth token on true
func (rh *RestHandler) Request(method string, payload interface{}, command string, auth bool) (*GeneralData, error) {

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

	var data []byte

	if auth {
		if rh.Token == "" {
			return nil, fmt.Errorf("token is empty. login first")
		}

		params := url.Values{}
		params.Add("token", rh.Token)
		params.Add("cmd", command)

		urlConcat = fmt.Sprintf("%s?%s", urlConcat, params.Encode())
	}

	data, err := json.Marshal([]interface{}{payload})

	if err != nil {
		return nil, err
	}

	reqUrl, err := url.Parse(urlConcat)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, reqUrl.String(), bytes.NewBuffer(data))

	if err != nil {
		return nil, err
	}

	req.Header = headers

	var client *http.Client

	// https://stackoverflow.com/questions/51845690/how-to-program-go-to-use-a-proxy-when-using-a-custom-transport
	// https://gist.github.com/ometa/71d23ed48c03c003f6e4910648612859
	if rh.Proxy != nil {

		tr := http.DefaultTransport.(*http.Transport).Clone()
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

		var proxyConcat string

		switch rh.Proxy.Type {
		case HTTP, HTTPS:
			if rh.Proxy.Auth != nil {
				proxyConcat = fmt.Sprintf("%s://%s:%s@%s:%d",
					rh.Proxy.Type,
					rh.Proxy.Auth.Username,
					rh.Proxy.Auth.Password,
					rh.Proxy.Host,
					rh.Port)
			} else {
				proxyConcat = fmt.Sprintf("%s://%s:%d",
					rh.Proxy.Type,
					rh.Proxy.Host,
					rh.Port)
			}

			proxyUrl, err := url.Parse(proxyConcat)

			if err != nil {
				return nil, err
			}

			tr.Proxy = http.ProxyURL(proxyUrl)
			client = &http.Client{Transport: tr}

			break
		case SOCKS5:
			proxyConcat = fmt.Sprintf("%s:%d",
				rh.Proxy.Host,
				rh.Port)

			var networkType string

			if rh.Proxy.Protocol != nil {
				networkType = rh.Proxy.Protocol.String()
			} else {
				networkType = "tcp"
			}

			dialer, err := proxy.SOCKS5(networkType, proxyConcat, nil, proxy.Direct)

			if contextDialer, ok := dialer.(proxy.ContextDialer); ok {
				tr.DialContext = contextDialer.DialContext
			} else {
				return nil, fmt.Errorf("failed to create socks5 dialer")
			}

			if err != nil {
				return nil, nil
			}

			break

		default:
			if rh.Proxy.Auth != nil {
				proxyConcat = fmt.Sprintf("http://%s:%s@%s:%d",
					rh.Proxy.Auth.Username,
					rh.Proxy.Auth.Password,
					rh.Proxy.Host,
					rh.Port)
			} else {
				proxyConcat = fmt.Sprintf("http://%s:%d",
					rh.Proxy.Host,
					rh.Port)
			}
		}

	} else {
		client = &http.Client{}
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	var result []*GeneralData
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return result[0], nil
}
