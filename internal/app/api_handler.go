package app

import (
	"github.com/ReolinkCameraAPI/reolinkapigo/internal/pkg/api"
	"github.com/ReolinkCameraAPI/reolinkapigo/internal/pkg/network/rest"
)

type ApiHandler struct {
	*api.AuthMixin
	*api.DeviceMixin
	*api.DisplayMixin
	*api.ImageMixin
	*api.RtspMixin
	*api.NetworkMixin
	*api.PtzMixin
	*api.RecordingMixin
	*api.SystemMixin
	*api.UserMixin
	*api.ZoomFocusMixin
	*rest.RestHandler
}

func NewApiHandler(username string, password string, host string, restOpts ...rest.OptionRestHandler) (*ApiHandler,
	error) {

	// create a new restHandler inside the apiHandler to manage all the rest network activity
	// such as injecting the token before a request is made.
	handler := rest.NewRestHandler(host, restOpts...)

	return &ApiHandler{
		&api.AuthMixin{
			Username: username,
			Password: password,
		},
		&api.DeviceMixin{},
		&api.DisplayMixin{},
		&api.ImageMixin{},
		&api.RtspMixin{
			Host:     host,
			Username: username,
			Password: password,
		},
		&api.NetworkMixin{},
		&api.PtzMixin{},
		&api.RecordingMixin{},
		&api.SystemMixin{},
		&api.UserMixin{},
		&api.ZoomFocusMixin{},
		handler,
	}, nil
}
