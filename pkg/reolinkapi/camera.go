package reolinkapi

import (
	"fmt"
	"github.com/ReolinkCameraAPI/reolinkapigo/internal/app"
	"github.com/ReolinkCameraAPI/reolinkapigo/internal/pkg/network/rest"
)

type Camera struct {
	*app.ApiHandler
}

type options struct {
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

func WithDeferLogin(deferLogin bool) OptionCamera {
	return deferLoginOption(deferLogin)
}

func WithNetworkOptions(networkOpts ...rest.OptionRestHandler) OptionCamera {
	return networkOption{networkOpts}
}

func NewCamera(username string, password string, ip string, opts ...OptionCamera) (
	*Camera, error) {

	options := options{
		deferLogin:  false,
		networkOpts: nil,
	}

	for _, o := range opts {
		o.apply(&options)
	}

	apiHandler, err := app.NewApiHandler(username, password, ip, options.networkOpts...)

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

	return &Camera{
		ApiHandler:  apiHandler,
	}, nil
}
