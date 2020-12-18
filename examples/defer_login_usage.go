package examples

import "github.com/ReolinkCameraAPI/reolinkapigo/pkg/reolinkapi"

func DeferLogin() {

	camera, err := reolinkapi.NewCamera("foo", "bar", "192.168.1.100",
		reolinkapi.WithDeferLogin(true),
	)

	if err != nil {
		panic(err)
	}

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
