package examples

import (
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/enum"
	"github.com/ReolinkCameraAPI/reolink-go-api/pkg"
	"gocv.io/x/gocv"
)

func RtspUsage() {
	// This can throw an error due to the API trying to authorise with the camera
	camera, err := pkg.NewCamera("foo", "bar", "192.168.1.100")

	if err != nil {
		panic(err)
	}

	// custom port 8554. The default will be used as 554
	streamPort := 8554
	rtspStream := camera.API.OpenRtspStream(&streamPort, enum.MAIN, nil)(camera.RestHandler)

	window := gocv.NewWindow("stream")

	for n := range rtspStream {
		if n.Err != nil {
			continue
		}
		window.IMShow(*n.Frame)
		window.WaitKey(1)
	}
}
