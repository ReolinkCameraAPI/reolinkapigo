package test

import (
	"fmt"
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/network"
	"testing"
	"time"
)

func TestNewRtspClient(t *testing.T) {

	rtspClient := network.NewRtspClient("rtsp://127.0.0.1:8544", network.RtspClientOptionDebug(true),
		network.RtspClientOptionRetryCount(1), network.RtspClientOptionTimeout(1*time.Second))

	defer rtspClient.OpenStream().Wait()

	for i := 0; i < 1; i++ {
		select {
		case pkt := <-rtspClient.Stream.Packet:
			if !pkt.IsKeyFrame {
				continue
			}

			fmt.Printf("Getting data %v", pkt.Data)
		}
		time.Sleep(5 * time.Second)
	}

}
