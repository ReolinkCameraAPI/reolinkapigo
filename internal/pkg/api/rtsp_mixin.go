package api

import (
	"fmt"
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/enum"
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/network"
	"gocv.io/x/gocv"
)

type RtspMixin struct {
}

type RtspStream struct {
	Frame *gocv.Mat
	Err   error
}

// Open an RTSP stream using GoCV (openCv 4)
// This function returns a channel containing the frames as it is received from the camera
// https://adaickalavan.github.io/portfolio/rtsp_video_streaming/#gsc.tab=0
func (rm *RtspMixin) OpenVideoStream(port *int, profile enum.RtspProfile,
	protocol *network.Protocol) func(handler *network.RestHandler) chan *RtspStream {
	return func(handler *network.RestHandler) chan *RtspStream {

		// creating unbuffered channel due to wanting frames in order
		// this will block the next frame from being accessible
		stream := make(chan *RtspStream)

		go func() {

			rtspUrl := fmt.Sprintf("%s:%s@%s", handler.Username, handler.Password, handler.Host)

			if port != nil {
				rtspUrl = fmt.Sprintf("%s:%d", rtspUrl, port)
			} else {
				rtspUrl = fmt.Sprintf("%s:%d", rtspUrl, 554)
			}

			rtspUrl = fmt.Sprintf("rtsp://%s//h264Preview_01_%s", rtspUrl, profile)

			capture, err := gocv.OpenVideoCaptureWithAPI(rtspUrl, gocv.VideoCaptureFFmpeg)

			if err != nil {
				stream <- &RtspStream{
					Frame: nil,
					Err:   err,
				}
			}

			frame := gocv.NewMat()

			for {

				if !capture.Read(&frame) {
					continue
				}

				stream <- &RtspStream{
					Frame: &frame,
					Err:   nil,
				}
			}

		}()

		return stream
	}
}
