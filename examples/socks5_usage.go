package examples

import (
	"github.com/ReolinkCameraAPI/reolinkapigo/internal/pkg/network/rest"
	"github.com/ReolinkCameraAPI/reolinkapigo/pkg/reolinkapi"
)

func Socks5Example() {

	// This can throw an error due to the API trying to authorise with the camera
	camera, err := reolinkapi.NewCamera("foo", "bar", "192.168.1.100",
		reolinkapi.WithNetworkOptions(
			rest.WithProxyScheme(rest.SOCKS5),
			rest.WithProxyHost("127.0.0.1"),
			rest.WithProxyPort(5942),
			rest.WithProxyUsername("foo"),
			rest.WithProxyPassword("bar"),
		))

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
