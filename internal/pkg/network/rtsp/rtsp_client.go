/*
Author: Benehiko
Modified Date: 2020/12/10
Contact: alanoterblanche@gmail.com | https://github.com/Benehiko

RtspClient is greatly inspired by github.com/deepch/RTSPtoWebRTC
RtspClient only wraps and simplifies their code to fit
inside this library.
*/
package rtsp

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
	UUID   string
	Stream *Stream
	*options
}

type options struct {
	username   string
	password   string
	timeout    time.Duration
	host       string
	port       int
	endpoint   string
	debug      bool
	retryCount int
}

type OptionRtspClient interface {
	apply(*options)
}

type timeoutOption int
type hostOption string
type portOption int
type debugOption bool
type retryOption int
type endpointOption string
type usernameOption string
type passwordOption string

func (t timeoutOption) apply(opts *options) {
	opts.timeout = time.Duration(t) * time.Second
}

func (h hostOption) apply(opts *options) {
	opts.host = string(h)
}

func (p portOption) apply(opts *options) {
	opts.port = int(p)
}

func (d debugOption) apply(opts *options) {
	opts.debug = bool(d)
}

func (r retryOption) apply(opts *options) {
	opts.retryCount = int(r)
}

func (e endpointOption) apply(opts *options) {
	opts.endpoint = string(e)
}

func (u usernameOption) apply(opts *options) {
	opts.username = string(u)
}

func (p passwordOption) apply(opts *options) {
	opts.password = string(p)
}

// Change the network timeout in seconds
func WithTimeout(timeout int) OptionRtspClient {
	return timeoutOption(timeout)
}

// Change the rtsp connection host (ip address or domain)
func WithHost(host string) OptionRtspClient {
	return hostOption(host)
}

func WithPort(port int) OptionRtspClient {
	return portOption(port)
}

func WithDebug(debug bool) OptionRtspClient {
	return debugOption(debug)
}

func WithRetry(retryCount int) OptionRtspClient {
	return retryOption(retryCount)
}

func WithEndpoint(endpoint string) OptionRtspClient {
	return endpointOption(endpoint)
}

func WithUsername(username string) OptionRtspClient {
	return usernameOption(username)
}

func WithPassword(password string) OptionRtspClient {
	return passwordOption(password)
}

// Create a new RTSP client
// Defaults:
// Timeout: 10 * time.Second
// Debug: false
// RetryCount: 5
// Port: 554
func NewRtspClient(host string, opts ...OptionRtspClient) *RtspClient {
	stream := &Stream{
		Codecs:  []av.CodecData{},
		Packets: make(map[string]chan av.Packet),
	}

	options := &options{
		username:   "",
		password:   "",
		timeout:    10,
		host:       host,
		port:       554,
		endpoint:   "",
		debug:      false,
		retryCount: 5,
	}

	for _, op := range opts {
		op.apply(options)
	}

	return &RtspClient{
		UUID:    uuid.UUIDv4(),
		Stream:  stream,
		options: options,
	}
}

func (rc *RtspClient) OpenStream() {
	log.Printf("Opening client stream: %s ", rc.UUID)

	url := "rtsp://"

	if rc.username != "" {
		url = fmt.Sprintf("%s%s:%s@", url, rc.username, rc.password)
	}

	url = fmt.Sprintf("%s%s:%d%s", url, rc.host, rc.port, rc.endpoint)

	for i := 0; i < rc.retryCount; i++ {
		session, err := rtsp.Dial(url)

		if err != nil {
			// Retry within 5 seconds
			fmt.Printf("err rtsp session: %s", err.Error())
			time.Sleep(5 * time.Second)
			continue
		}

		session.RtpKeepAliveTimeout = rc.timeout

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
