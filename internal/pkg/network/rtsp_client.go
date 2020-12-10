/*
RtspClient is greatly inspired by github.com/deepch/RTSPtoWebRTC
RtspClient only wraps and simplifies their code to fit
inside this library.
*/
package network

import (
	"github.com/deepch/vdk/av"
	"github.com/deepch/vdk/format/rtsp"
	"sync"
	"time"
)

// stream contains all the necessary information of
// the stream.
// stream packet contains audio and video
type stream struct {
	Codecs []av.CodecData
	Packet chan av.Packet
}

type RtspClient struct {
	Timeout    time.Duration
	URL        string
	Debug      bool
	RetryCount int
	Status     *RtspClientStatus
	Stream     *stream
}

type RtspClientStatus struct {
	IsOpen chan bool
}

type OptionRtspClient func(*RtspClient)

// Create a new RTSP client
// Defaults:
// Timeout: 10 * time.Second
// Debug: false
// RetryCount: 5
func NewRtspClient(url string, rtspOptions ...OptionRtspClient) *RtspClient {
	stream := &stream{
		Codecs: nil,
		Packet: make(chan av.Packet),
	}

	isOpen := make(chan bool)
	isOpen <- false

	status := &RtspClientStatus{
		IsOpen: isOpen,
	}

	rtspClient := &RtspClient{
		Timeout:    10 * time.Second,
		URL:        url,
		Debug:      false,
		RetryCount: 5,
		Status:     status,
		Stream:     stream,
	}

	for _, op := range rtspOptions {
		op(rtspClient)
	}

	return rtspClient
}

func (rc *RtspClient) OpenStream() *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	// Spawn on separate thread
	go func() {
		defer wg.Done()
		for i := 0; i < rc.RetryCount; i++ {
			session, err := rtsp.Dial(rc.URL)

			if err != nil {
				// Retry within 5 seconds
				time.Sleep(5 * time.Second)
				continue
			}

			session.RtpKeepAliveTimeout = rc.Timeout

			codec, err := session.Streams()
			if err != nil {
				// Retry within 5 seconds
				time.Sleep(5 * time.Second)
				continue
			}

			rc.Stream.Codecs = codec

			rc.Status.IsOpen <- true

			for <-rc.Status.IsOpen {
				pkt, err := session.ReadPacket()

				if err != nil {
					rc.Status.IsOpen <- false
					break
				}

				rc.Stream.Packet <- pkt
			}

			err = session.Close()

			if err != nil {
				// Retrying in 5 seconds
				time.Sleep(5 * time.Second)
			}

		}
	}()
	return wg
}

func (rc *RtspClient) StopStream() {
	rc.Status.IsOpen <- false
}

func RtspClientOptionTimeout(timeout time.Duration) OptionRtspClient {
	return func(client *RtspClient) {
		client.Timeout = timeout
	}
}

func RtspClientOptionDebug(debug bool) OptionRtspClient {
	return func(client *RtspClient) {
		client.Debug = debug
	}
}

func RtspClientOptionRetryCount(retryCount int) OptionRtspClient {
	return func(client *RtspClient) {
		client.RetryCount = retryCount
	}
}
