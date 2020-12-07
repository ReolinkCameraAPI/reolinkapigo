package examples

import "github.com/ReolinkCameraAPI/reolink-go-api/pkg"

func BasicUsage() {

	// This can throw an error due to the API trying to authorise with the camera
	camera, err := pkg.NewCamera("foo", "bar", "192.168.1.100")

	if err != nil {
		panic(err)
	}

	// Call your camera api here and pass the camera restHandler to the function
	ok, err := camera.API.FormatHdd(0)(camera.RestHandler)

	if err != nil {
		panic(err)
	}

	if ok {
		print("Format OK")
	}
}
