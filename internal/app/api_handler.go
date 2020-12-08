package app

import (
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/api"
)

type ApiHandler struct {
	*api.DeviceMixin
	*api.DisplayMixin
	*api.ImageMixin
	*api.AuthMixin
	*api.RtspMixin
	*api.NetworkMixin
}

func NewApiHandler() (*ApiHandler, error) {
	return &ApiHandler{
		&api.DeviceMixin{},
		&api.DisplayMixin{},
		&api.ImageMixin{},
		&api.AuthMixin{},
		&api.RtspMixin{},
		&api.NetworkMixin{},
	}, nil
}
