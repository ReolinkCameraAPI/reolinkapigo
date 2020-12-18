package app

import (
	"github.com/ReolinkCameraAPI/reolinkapigo/internal/pkg/api"
	"github.com/ReolinkCameraAPI/reolinkapigo/internal/pkg/network"
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
	*network.RestHandler
}

func NewApiHandler(username string, password string, host string, restOpts ...network.OptionRestHandler) (*ApiHandler,
	error) {

	handler := network.NewRestHandler(host, restOpts...)

	return &ApiHandler{
		&api.AuthMixin{
			Username: username,
			Password: password,
		},
		&api.DeviceMixin{},
		&api.DisplayMixin{},
		&api.ImageMixin{},
		&api.RtspMixin{
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
