package reolinkapi

import (
	"fmt"
	"github.com/ReolinkCameraAPI/reolinkapigo/internal/app"
	"github.com/ReolinkCameraAPI/reolinkapigo/internal/pkg/network/rest"
	"golang.org/x/net/context"
	"time"
)

type Camera struct {
	*app.ApiHandler
}

type options struct {
	username    string
	password    string
	deferLogin  bool
	networkOpts []rest.OptionRestHandler
}

type OptionCamera interface {
	apply(*options)
}

type deferLoginOption bool

func (d deferLoginOption) apply(opts *options) {
	opts.deferLogin = bool(d)
}

type networkOption struct {
	networkOpts []rest.OptionRestHandler
}

func (n networkOption) apply(opts *options) {
	opts.networkOpts = n.networkOpts
}

type usernameOption string

func (u usernameOption) apply(opts *options) {
	opts.username = string(u)
}

type passwordOption string

func (p passwordOption) apply(opts *options) {
	opts.password = string(p)
}

func WithDeferLogin(deferLogin bool) OptionCamera {
	return deferLoginOption(deferLogin)
}

func WithNetworkOptions(networkOpts ...rest.OptionRestHandler) OptionCamera {
	return networkOption{networkOpts}
}

func WithUsername(username string) OptionCamera {
	return usernameOption(username)
}

func WithPassword(password string) OptionCamera {
	return passwordOption(password)
}

// Create a new camera object
// IP is required. Username and Password will fallback to camera defaults.
// To change the network options such as connecting to a camera behind a proxy, pass the networkOpts parameter
// Defaults:
// Username: "admin"
// Password: ""
// deferLogin: false
// networkOpts: nil
func NewCamera(ip string, opts ...OptionCamera) (
	*Camera, error) {

	options := options{
		deferLogin:  false,
		networkOpts: nil,
		username:    "admin",
		password:    "",
	}

	for _, o := range opts {
		o.apply(&options)
	}

	apiHandler, err := app.NewApiHandler(options.username, options.password, ip, options.networkOpts...)

	if err != nil {
		return nil, err
	}

	// login immediately if deferLogin is false
	// else leave it up to the user to decide when to log in.
	if !options.deferLogin {
		// pass the restHandler object to the Login function
		ok, err := apiHandler.Login()(apiHandler.RestHandler)

		if err != nil {
			return nil, err
		}

		if !ok {
			return nil, fmt.Errorf("login unsuccessful")
		}
	}

	camera := &Camera{
		ApiHandler: apiHandler,
	}

	return camera, nil
}

// Auto refresh
func (c *Camera) AutoRefreshToken(ctx context.Context) chan error {
	err := make(chan error)
	go func() {
		for {
			if c.IsLoggedIn() {
				ok, e := c.IsTokenValid()
				if err != nil {
					err <- e
					return
				}
				if !ok {
					_, e := c.Login()(c.RestHandler)

					if e != nil {
						err <- e
					}
				}
			}

			select {
			case <-ctx.Done():
				return // terminate goroutine
			default: // avoid blocking
			}
			time.Sleep(time.Minute * 5)
		}
	}()

	return err
}
