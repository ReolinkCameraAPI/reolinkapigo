package examples

import (
	"context"
	"fmt"
	"github.com/ReolinkCameraAPI/reolinkapigo/pkg/reolinkapi"
)

func DeferLogin() {

	camera, err := reolinkapi.NewCamera("192.168.1.100",
		reolinkapi.WithUsername("foo"),
		reolinkapi.WithPassword("bar"),
		reolinkapi.WithDeferLogin(true),
	)

	if err != nil {
		panic(err)
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	// this will attempt to auto refresh the token automatically.
	// if some sort of network error or authentication error occurs
	// it will return an error in its channel.
	// Don't worry about logging in before calling this.
	// it will only attempt to refresh the token once
	// there has been a previous token set
	errChan := camera.AutoRefreshToken(ctx)

	go func() {
		select {
		case err := <-errChan:
			fmt.Printf("AutoRefreshToken got an error: %s", err.Error())
		}
	}()

	// before you can do anything call login since the camera won't have any token
	_, err = camera.Login()(camera.RestHandler)

	if err != nil {
		panic(err)
	}

	// Call your camera api here and pass the camera restHandler to the function
	ok, err := camera.FormatHdd(0)(camera.RestHandler)

	if err != nil {
		panic(err)
	}

	if ok {
		print("Format OK")
	}

}
