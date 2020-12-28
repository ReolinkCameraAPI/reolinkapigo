package examples

import (
	"context"
	"fmt"
	"github.com/ReolinkCameraAPI/reolinkapigo/pkg/reolinkapi"
)

func BasicUsage() {

	// This can throw an error due to the API trying to authorise with the camera
	camera, err := reolinkapi.NewCamera("192.168.1.100",
		reolinkapi.WithUsername("foo"),
		reolinkapi.WithPassword("bar"))

	if err != nil {
		panic(err)
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	// this will attempt to auto refresh the token automatically.
	// if some sort of network error or authentication error occurs
	// it will return an error in its channel.
	errChan := camera.AutoRefreshToken(ctx)

	go func() {
		select {
		case err := <-errChan:
			fmt.Printf("AutoRefreshToken got an error: %s", err.Error())
		}
	}()

	// Call your camera api here and pass the camera restHandler to the function
	ok, err := camera.FormatHdd(0)(camera.RestHandler)

	if err != nil {
		panic(err)
	}

	if ok {
		print("Format OK")
	}
}
