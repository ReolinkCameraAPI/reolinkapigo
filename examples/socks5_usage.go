package examples

import (
	"github.com/ReolinkCameraAPI/reolinkapigo/internal/pkg/network"
	"github.com/ReolinkCameraAPI/reolinkapigo/pkg/reolinkapi"
)

func Socks5Example() {

	// This can throw an error due to the API trying to authorise with the camera
	camera, err := reolinkapi.NewCamera("foo", "bar", "192.168.1.100",
		reolinkapi.WithNetworkOptions(
			network.WithProxyScheme(network.SOCKS5),
			network.WithProxyHost("127.0.0.1"),
			network.WithProxyPort(5942),
			network.WithProxyUsername("foo"),
			network.WithProxyPassword("bar"),
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
