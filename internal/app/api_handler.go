package app

import (
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/api"
)

type ApiHandler struct {
	*api.DeviceMixin
	*api.DisplayMixin
	*api.ImageMixin
	*api.AuthMixin
}

func NewApiHandler(username string, password string) (*ApiHandler, error) {
	authMixin := &api.AuthMixin{
		Username: username,
		Password: password,
	}

	return &ApiHandler{
		&api.DeviceMixin{},
		&api.DisplayMixin{},
		&api.ImageMixin{},
		authMixin,
	}, nil
}
