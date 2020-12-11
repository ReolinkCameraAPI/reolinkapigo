package test

import (
	"bytes"
	"fmt"
	"github.com/ReolinkCameraAPI/reolink-go-api/examples"
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/network"
	"image"
	"testing"
	"time"
)

func TestNewRtspClient(t *testing.T) {

	rtspClient := network.NewRtspClient("rtsp://localhost:8554/mystream", network.RtspClientOptionDebug(true),
		network.RtspClientOptionRetryCount(1), network.RtspClientOptionTimeout(1*time.Second))

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
