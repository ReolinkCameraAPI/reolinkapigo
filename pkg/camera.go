package pkg

import (
	"fmt"
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/app"
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/network"
)

type Camera struct {
	API         *app.ApiHandler
	RestHandler *network.RestHandler
}

func NewCamera(username string, password string, ip string, restOptions ...network.OptionRestHandler) (
	*Camera, error) {
	restHandler := network.NewRestHandler(ip, restOptions...)

	restHandler.SetUsernamePassword(username, password)

	apiHandler, err := app.NewApiHandler()

	if err != nil {
		return nil, err
	}

	// pass the restHandler object to the Login function
	ok, err := apiHandler.Login()(restHandler)

	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, fmt.Errorf("login unsuccessful")
	}

	return &Camera{
		API:         apiHandler,
		RestHandler: restHandler,
	}, nil
}
