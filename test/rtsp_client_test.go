package test

import (
	"bytes"
	"fmt"
	"github.com/ReolinkCameraAPI/reolinkapigo/examples"
	"github.com/ReolinkCameraAPI/reolinkapigo/internal/pkg/network/rtsp"
	"image"
	"testing"
)

func TestNewRtspClient(t *testing.T) {

	rtspClient := rtsp.NewRtspClient("rtsp://localhost:8554/mystream")

	go rtspClient.OpenStream()

	println("got stream open and awaiting response")

	for key, _ := range rtspClient.Stream.Packets {
		go func(packetUUID string) {
			for {
				select {
				case pkt := <-rtspClient.Stream.Packets[packetUUID]:
					fmt.Printf("Getting data...")
					reader := bytes.NewReader(pkt.Data)

					_, _, err := image.Decode(reader)

					if err != nil {
						continue
					}
					// Do something with image here
				}
			}
		}(key)
	}

}

func TestNewRtspView(t *testing.T) {
	examples.RtspUsage()
}
