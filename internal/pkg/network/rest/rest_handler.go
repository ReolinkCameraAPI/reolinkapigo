package rest

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/net/proxy"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var (
	headers = map[string][]string{
		"Content-Class": {"application/json"},
		"Accept":        {"application/json"},
	}
)

type optionsProxy struct {
	scheme   Scheme
	protocol Protocol
	host     string
	port     int
	username string
	password string
}

type options struct {
	host     string
	port     int
	endpoint string
	scheme   Scheme
	token    string
	proxy    *optionsProxy
}

type OptionRestHandler interface {
	apply(*options)
}

type portOption int
type endpointOption string
type schemeOption struct {
	Scheme
}

type proxyScheme struct {
	Scheme
}
type proxyProtocol struct {
	Protocol
}
type proxyHost string
type proxyPort int
type proxyUsername string
type proxyPassword string

func (p portOption) apply(opts *options) {
	opts.port = int(p)
}

func (e endpointOption) apply(opts *options) {
	opts.endpoint = string(e)
}

func (s schemeOption) apply(opts *options) {
	opts.scheme = s.Scheme
}

func (p proxyScheme) apply(opts *options) {
	opts.proxy.scheme = p.Scheme
}

func (p proxyProtocol) apply(opts *options) {
	opts.proxy.protocol = p.Protocol
}

func (p proxyHost) apply(opts *options) {
	opts.proxy.host = string(p)
}

func (p proxyPort) apply(opts *options) {
	opts.proxy.port = int(p)
}

func (p proxyUsername) apply(opts *options) {
	opts.proxy.username = string(p)
}

func (p proxyPassword) apply(opts *options) {
	opts.proxy.password = string(p)
}

// Change the default port to a custom port.
// Default is unset due to http being the default protocol
func WithPort(port int) OptionRestHandler {
	return portOption(port)
}

// Change the default endpoint to a custom endpoint
// Default is "/cgi-bin/api.cgi"
// If for some reason the camera you are using is different, one can update it here.
func WithEndpoint(endpoint string) OptionRestHandler {
	return endpointOption(endpoint)
}

// Change the default scheme from HTTP to HTTPS or SOCKS5
func WithScheme(scheme Scheme) OptionRestHandler {
	return schemeOption{scheme}
}

func WithProxyProtocol(protocol Protocol) OptionRestHandler {
	return proxyProtocol{protocol}
}

// Add a username to the proxy configuration
func WithProxyUsername(username string) OptionRestHandler {
	return proxyUsername(username)
}

// Add a password to the proxy configuration
func WithProxyPassword(password string) OptionRestHandler {
	return proxyPassword(password)
}

// Change the default scheme from HTTP to HTTPS or SOCKS5
func WithProxyScheme(scheme Scheme) OptionRestHandler {
	return proxyScheme{scheme}
}

// Add a proxy host configuration
func WithProxyHost(host string) OptionRestHandler {
	return proxyHost(host)
}

// Add a proxy port configuration
func WithProxyPort(port int) OptionRestHandler {
	return proxyPort(port)
}

type RestHandler struct {
	*options
}

// Create a new RestHandler object with optional argument using Variadic options pattern for customisation
// Refer to the RestHandlerOption<option_name> functions
// RestHandler is used to wrap the http package and give a cleaner more defined scope which the person
// implementing the library will have full control over.
// https://stackoverflow.com/a/26326418
func NewRestHandler(host string, opts ...OptionRestHandler) *RestHandler {
	options := &options{
		host:     host,
		port:     0,
		endpoint: "cgi-bin/api.cgi",
		scheme:   HTTP,
		token:    "",
		proxy: &optionsProxy{
			scheme:   HTTP,
			protocol: PROTOCOL_TCP,
			host:     "",
			port:     0,
			username: "",
			password: "",
		},
	}

	for _, op := range opts {
		op.apply(options)
	}

	return &RestHandler{options}
}

// Do the http request
// endpoint: the trailing part of the URL after the port.
// method: GET or POST
// payload: the json data
// auth: alters the request to include auth token on true
func (rh *RestHandler) Request(method string, payload interface{}, command string) (*GeneralData, error) {

	params := url.Values{}
	params.Add("cmd", command)

	respBody, err := rh.RequestRaw(method, payload, params)

	if err != nil {
		return nil, err
	}

	var result []*GeneralData

	err = json.Unmarshal(respBody, &result)

	if err != nil {
		return nil, err
	}

	return result[0], nil
}

func (rh *RestHandler) RequestRaw(method string, payload interface{}, params url.Values) ([]byte, error) {
	var urlConcat string
	if rh.port > 0 {
		urlConcat = fmt.Sprintf("%s:%d/%s", rh.host, rh.port, rh.endpoint)
	} else {
		urlConcat = fmt.Sprintf("%s/%s", rh.host, rh.endpoint)
	}

	urlConcat = fmt.Sprintf("%s://%s", rh.scheme.String(), urlConcat)

	params.Add("token", rh.token)

	urlConcat = fmt.Sprintf("%s?%s", urlConcat, params.Encode())

	var data []byte

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
	if rh.proxy.host != "" {

		tr := http.DefaultTransport.(*http.Transport).Clone()
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

		var proxyConcat string

		switch rh.proxy.scheme {
		case HTTP, HTTPS:
			proxyConcat = fmt.Sprintf("%s://", rh.proxy.scheme)

			if rh.proxy.username != "" {
				proxyConcat = fmt.Sprintf("%s%s:%s@%s:%d",
					proxyConcat,
					rh.proxy.username,
					rh.proxy.password,
					rh.proxy.host,
					rh.proxy.port)

			} else {
				proxyConcat = fmt.Sprintf("%s%s:%d",
					proxyConcat,
					rh.proxy.host,
					rh.proxy.port)
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
				rh.proxy.host,
				rh.proxy.port)

			networkType := rh.proxy.protocol.String()

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
			proxyConcat = fmt.Sprintf("%s://", rh.proxy.scheme)

			if rh.proxy.username != "" {
				proxyConcat = fmt.Sprintf("%s%s:%s@%s:%d",
					proxyConcat,
					rh.proxy.username,
					rh.proxy.password,
					rh.proxy.host,
					rh.proxy.port)

			} else {
				proxyConcat = fmt.Sprintf("%s%s:%d",
					proxyConcat,
					rh.proxy.host,
					rh.proxy.port)
			}
		}

	} else {
		client = &http.Client{}
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

// Set the current token
func (rh *RestHandler) SetToken(token string) {
	rh.token = token
}

// Get the current token
func (rh *RestHandler) GetToken() string {
	return rh.token
}

// Check if there is a token
func (rh *RestHandler) IsLoggedIn() bool {
	if rh.token != "" {
		return true
	}
	return false
}

// Check if the token is valid
// Will return true if valid
func (rh *RestHandler) IsTokenValid() (bool, error) {
	claims := jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(rh.token, claims, func(token *jwt.Token) (interface{}, error) {
		return token, nil
	})

	if err != nil {
		return false, err
	}

	if !token.Valid {
		return false, nil
	}

	tm := time.Unix(claims.ExpiresAt, 0)
	remainder := tm.Sub(time.Now())

	remaining := remainder.Seconds()

	if remaining <= 0 {
		return false, nil
	}

	return true, nil
}
