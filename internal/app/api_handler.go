package app

import (
	"fmt"
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/api"
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/network"
)

type ApiHandler struct {
	*network.RestHandler
	*api.DeviceMixin
	*api.DisplayMixin
	*api.ImageMixin
	*api.AuthMixin
}

func NewApiHandler(username string, password string, restHandler *network.RestHandler) (*ApiHandler, error) {
	authMixin := &api.AuthMixin{
		Username: username,
		Password: password,
	}

	// pass the restHandler object to the Login function
	ok, err := authMixin.Login()(restHandler)

	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, fmt.Errorf("login unsuccessful")
	}

	return &ApiHandler{
		restHandler,
		&api.DeviceMixin{},
		&api.DisplayMixin{},
		&api.ImageMixin{},
		authMixin,
	}, nil
}
