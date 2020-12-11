/*
Author: Benehiko
Modified Date: 2020/12/10
Contact: alanoterblanche@gmail.com | https://github.com/Benehiko

RtspClient is greatly inspired by github.com/deepch/RTSPtoWebRTC
RtspClient only wraps and simplifies their code to fit
inside this library.
*/
package network

import (
	"fmt"
	"github.com/deepch/vdk/av"
	"github.com/deepch/vdk/format/rtsp"
	"github.com/dgryski/trifles/uuid"
	"log"
	"time"
)

// stream contains all the necessary information of
// the stream.
// stream packet contains audio and video
type Stream struct {
	Codecs  []av.CodecData
	Packets map[string]chan av.Packet
}

type RtspClient struct {
	UUID       string
	Timeout    time.Duration
	URL        string
	Debug      bool
	RetryCount int
	Stream     *Stream
}

type OptionRtspClient func(*RtspClient)

// Create a new RTSP client
// Defaults:
// Timeout: 10 * time.Second
// Debug: false
// RetryCount: 5
func NewRtspClient(url string, rtspOptions ...OptionRtspClient) *RtspClient {
	stream := &Stream{
		Codecs:  []av.CodecData{},
		Packets: make(map[string]chan av.Packet),
	}

	rtspClient := &RtspClient{
		UUID:       uuid.UUIDv4(),
		Timeout:    10 * time.Second,
		URL:        url,
		Debug:      false,
		RetryCount: 5,
		Stream:     stream,
	}

	for _, op := range rtspOptions {
		op(rtspClient)
	}

	return rtspClient
}

func (rc *RtspClient) OpenStream() {
	log.Printf("Opening client stream: %s ", rc.UUID)

	for i := 0; i < rc.RetryCount; i++ {
		session, err := rtsp.Dial(rc.URL)

		if err != nil {
			// Retry within 5 seconds
			fmt.Printf("err rtsp session: %s", err.Error())
			time.Sleep(5 * time.Second)
			continue
		}

		session.RtpKeepAliveTimeout = rc.Timeout

		codecs, err := session.Streams()
		if err != nil {
			fmt.Printf("err rtsp codec: %s", err.Error())
			// Retry within 5 seconds
			time.Sleep(5 * time.Second)
			continue
		}

		fmt.Printf("Codecs type %s", codecs[1].Type())

		rc.Stream.Codecs = codecs

		isOpen := true

		for isOpen {
			pkt, err := session.ReadPacket()

			if err != nil {
				fmt.Printf("err rtsp packet: %s", err.Error())
				isOpen = false
				break
			}

			// This is mutable from the calling method
			// Will increase as more packets are created
			// Check web_rtc_client.go -> line 213
			// | packetUUID := uuid.UUIDv4()
			// | wrtc.Stream.Packets[packetUUID] = rtspPacketChannel
			for _, v := range rc.Stream.Packets {
				if len(v) < cap(v) {
					v <- pkt
				}
			}

		}

		err = session.Close()

		if err != nil {
			log.Println("Session close error.", err)
			// Retrying in 5 seconds
			time.Sleep(5 * time.Second)
		}

	}
}

/*func (rc *RtspClient) StopStream() {
	rc.Status.IsOpen <- false
}*/

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
