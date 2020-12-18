/*
Author: Benehiko
Modified Date: 2020/12/10
Contact: alanoterblanche@gmail.com | https://github.com/Benehiko

WebRtcClient is greatly inspired by github.com/deepch/RTSPtoWebRTC
WebRtcClient only wraps and simplifies their code to fit
inside this library.
*/
package rtsp

import (
	"encoding/base64"
	"github.com/deepch/vdk/av"
	"github.com/deepch/vdk/codec/h264parser"
	"github.com/dgryski/trifles/uuid"
	"github.com/pion/webrtc/v2"
	"github.com/pion/webrtc/v2/pkg/media"
	"log"
	"math/rand"
	"strings"
	"time"
)

type WebRtcClient struct {
	Ready chan bool
	SDP   chan string
	RTSP  *RtspClient
}

func NewWebRtcClient(rtspClient *RtspClient) *WebRtcClient {
	return &WebRtcClient{
		Ready: make(chan bool),
		SDP:   make(chan string),
		RTSP:  rtspClient,
	}
}

// Open a WebRTC Stream from the RTSP Stream
// Credit for most of this code goes to deepch https://github.com/deepch/RTSPtoWebRTC
func (wrtc *WebRtcClient) OpenWebRtcStream(sdpData string) {
	if wrtc != nil {

		if wrtc.RTSP.Stream.Codecs == nil {
			log.Println("Codec error")
			return
		}

		sps := wrtc.RTSP.Stream.Codecs[0].(h264parser.CodecData).SPS()
		pps := wrtc.RTSP.Stream.Codecs[0].(h264parser.CodecData).PPS()

		// Receive the remote SDP as Base64
		sd, err := base64.StdEncoding.DecodeString(sdpData)

		if err != nil {
			log.Println("DecodeString error", err)
			return
		}

		// Create MediaEngine
		mediaEngine := webrtc.MediaEngine{}
		offer := webrtc.SessionDescription{
			Type: webrtc.SDPTypeOffer,
			SDP:  string(sd),
		}

		err = mediaEngine.PopulateFromSDP(offer)
		if err != nil {
			log.Println("PopulateFromSDP error", err)
			return
		}

		var payloadType uint8
		for _, videoCodec := range mediaEngine.GetCodecsByKind(webrtc.RTPCodecTypeVideo) {
			if videoCodec.Name == "H264" && strings.Contains(videoCodec.SDPFmtpLine, "packetization-mode=1") {
				payloadType = videoCodec.PayloadType
				break
			}
		}

		if payloadType == 0 {
			log.Println("Remote peer does not support H264")
			return
		}

		if payloadType != 126 {
			log.Println("Video might not work with codec", payloadType)
		}

		log.Println("Work payloadType", payloadType)
		webRtcApi := webrtc.NewAPI(webrtc.WithMediaEngine(mediaEngine))

		// TODO: Need to investigate this
		peerConnection, err := webRtcApi.NewPeerConnection(webrtc.Configuration{
			ICEServers: []webrtc.ICEServer{
				{
					URLs: []string{"stun:stun.l.google.com:19302"},
				},
			},
		})

		if err != nil {
			log.Println("NewPeerConnection error", err)
			return
		}

		// Add KeepAlive timer
		peerConnTimer := time.NewTimer(time.Second * 2)
		peerConnection.OnDataChannel(func(d *webrtc.DataChannel) {
			// Register text message handling
			d.OnMessage(func(msg webrtc.DataChannelMessage) {
				peerConnTimer.Reset(2 * time.Second)
			})
		})

		// Add Video Track
		videoTrack, err := peerConnection.NewTrack(
			payloadType,
			rand.Uint32(),
			"video",
			wrtc.RTSP.UUID+"_pion")

		if err != nil {
			log.Println("Could not create VideoTrack", err)
			return
		}

		_, err = peerConnection.AddTransceiverFromTrack(videoTrack,
			webrtc.RtpTransceiverInit{
				Direction: webrtc.RTPTransceiverDirectionSendonly,
			},
		)

		if err != nil {
			log.Println("AddTransceiverFromTrack error", err)
			return
		}

		_, err = peerConnection.AddTrack(videoTrack)
		if err != nil {
			log.Println("AddTrack error", err)
			return
		}

		// Add Audio Track
		var audioTrack *webrtc.Track
		if len(wrtc.RTSP.Stream.Codecs) > 1 && (wrtc.RTSP.Stream.Codecs[1].Type() == av.PCM_ALAW || wrtc.
			RTSP.Stream.Codecs[1].Type() == av.PCM_MULAW) {

			switch wrtc.RTSP.Stream.Codecs[1].Type() {
			case av.PCM_ALAW:
				audioTrack, err = peerConnection.NewTrack(webrtc.DefaultPayloadTypePCMA, rand.Uint32(), "audio",
					wrtc.RTSP.UUID+"audio")
			case av.PCM_MULAW:
				audioTrack, err = peerConnection.NewTrack(webrtc.DefaultPayloadTypePCMU, rand.Uint32(), "audio",
					wrtc.RTSP.UUID+"audio")
			}

			if err != nil {
				log.Println(err)
				return
			}

			_, err = peerConnection.AddTransceiverFromTrack(audioTrack,
				webrtc.RtpTransceiverInit{
					Direction: webrtc.RTPTransceiverDirectionSendonly,
				},
			)

			if err != nil {
				log.Println("AddTransceiverFromTrack error", err)
				return
			}

			_, err = peerConnection.AddTrack(audioTrack)

			if err != nil {
				log.Println(err)
				return
			}

		}

		if err = peerConnection.SetRemoteDescription(offer); err != nil {
			log.Println("SetRemoteDescription error", err, offer.SDP)
			return
		}

		peerAnswer, err := peerConnection.CreateAnswer(nil)

		if err != nil {
			log.Println("CreateAnswer error", err)
			return
		}

		if err = peerConnection.SetLocalDescription(peerAnswer); err != nil {
			log.Println("SetLocalDescription error", err)
			return
		}

		wrtc.SDP <- base64.StdEncoding.EncodeToString([]byte(peerAnswer.SDP))

		/*if err != nil {
			log.Println("Writing SDP error", err)
			return
		}*/

		peerConnectionControl := make(chan bool, 10)

		peerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
			log.Printf("Connection State has changed %s\n", connectionState.String())

			if connectionState != webrtc.ICEConnectionStateConnected {
				log.Println("Client Close Exit")

				err = peerConnection.Close()

				if err != nil {
					log.Println("peerConnection Close error", err)
				}

				peerConnectionControl <- true
				return
			}

			if connectionState == webrtc.ICEConnectionStateConnected {
				go func() {
					rtspPacketChannel := make(chan av.Packet, 100)

					// Create a new packet identifier and it's own buffer
					// This will mutate the current streams packets.
					packetUUID := uuid.UUIDv4()
					wrtc.RTSP.Stream.Packets[packetUUID] = rtspPacketChannel

					// Defer the deletion of the packet channel
					defer func() {
						defer delete(wrtc.RTSP.Stream.Packets, packetUUID)
					}()

					var Vpre time.Duration
					// reset timeout for client
					peerConnTimer.Reset(time.Second * 5)
					var start bool
					for {
						select {
						case <-peerConnTimer.C:
							log.Println("Client close Keep-Alive Timer")
							peerConnection.Close()
						case <-peerConnectionControl:
							return
						case pkt := <-rtspPacketChannel:
							if pkt.IsKeyFrame {
								// is a keyframe indicating it should start
								start = true
								pkt.Data = append([]byte("\000\000\001"+string(sps)+"\000\000\001"+string(
									pps)+"\000\000\001"), pkt.Data[4:]...)
							}else if start {
								// not a keyframe but has started
								pkt.Data = pkt.Data[4:]
							} else {
								// not a keyframe and not started
								continue
							}

							var Vts time.Duration
							if pkt.Idx == 0 && videoTrack != nil {
								if Vpre != 0 {
									Vts = pkt.Time - Vpre
								}
								samples := uint32(90000 / 1000 * Vts.Milliseconds())
								err := videoTrack.WriteSample(media.Sample{Data: pkt.Data, Samples: samples})
								if err != nil {
									return
								}
								Vpre = pkt.Time
							} else if pkt.Idx == 1 && audioTrack != nil {
								err := audioTrack.WriteSample(media.Sample{Data: pkt.Data,
									Samples: uint32(len(pkt.Data))})
								if err != nil {
									return
								}
							}
						}
					}

				}()
			}
		})
		return
	}
}
