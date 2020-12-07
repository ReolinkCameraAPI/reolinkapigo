package examples

import (
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/network"
	"github.com/ReolinkCameraAPI/reolink-go-api/pkg"
)

func Socks5Example() {

	// Setting the RestHandler to do all requests over HTTPS
	httpsOptions := network.RestHandlerOptionHttp(true)

	// Setting the RestHandler to proxy requests through SOCKS
	// the default protocol is tcp
	protocol := network.UDP
	socksOptions := network.RestHandlerOptionProxy(network.SOCKS5, "127.0.0.1", 5942, nil, &protocol)

	// This can throw an error due to the API trying to authorise with the camera
	camera, err := pkg.NewCamera("foo", "bar", "192.168.1.100", socksOptions, httpsOptions)

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
