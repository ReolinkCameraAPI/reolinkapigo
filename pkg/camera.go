package pkg

import (
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/app"
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/network"
)

type Camera struct {
	API *app.ApiHandler
}

func NewCamera(username string, password string, ip string, restOptions ...func(handler *network.RestHandler) error) (
	*Camera, error) {
	restHandler, err := network.NewRestHandler(ip, restOptions...)

	if err != nil {
		return nil, err
	}

	apiHandler, err := app.NewApiHandler(username, password, restHandler)

	if err != nil {
		return nil, err
	}

	return &Camera{
		API: apiHandler,
	}, nil
}
