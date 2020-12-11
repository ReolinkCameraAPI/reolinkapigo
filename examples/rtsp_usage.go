package examples

import (
	"encoding/json"
	"fmt"
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/network"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sort"
	"sync"
	"time"
)

type Stream struct {
	RtspClients []*network.RtspClient
}

// RTSP stream to WebRTC
// Code for RTSP to WebRTC by https://github.com/deepch/RTSPtoWebRTC
func RtspUsage() {

	rtspStreamUrls := []string{
		"rtsp://localhost:8554/mystream",
	}

	rtspClients := []*network.RtspClient{}

	// Create all the RTSP clients
	for _, u := range rtspStreamUrls {
		rtspClient := network.NewRtspClient(u, network.RtspClientOptionDebug(true),
			network.RtspClientOptionRetryCount(1), network.RtspClientOptionTimeout(1*time.Second))

		fmt.Printf("Opening stream %s...\n", u)

		rtspClients = append(rtspClients, rtspClient)
	}

	streams := &Stream{RtspClients: rtspClients}

	// Open all the streams
	for _, rc := range rtspClients {
		go rc.OpenStream()
	}

	// Start the web service
	defer streams.webServer().Wait()
}

// Handle the stream
// Sends WebRTC stream sourced from RTSP stream
func (s *Stream) handlerStream(c *gin.Context) {
	fmt.Println("New Client connected")

	c.Header("Access-Control-Allow-Origin", "*")
	sdpData := c.PostForm("data")
	streamUUID := c.PostForm("streamUUID")

	var currentRtsp *network.RtspClient

	for _, rc := range s.RtspClients {
		if rc.UUID == streamUUID {
			currentRtsp = rc
			break
		}
	}

	if currentRtsp != nil {
		webRtcStream := network.NewWebRtcClient(currentRtsp)

		// TODO:
		go webRtcStream.OpenWebRtcStream(nil, sdpData)

		for {
			select {
			case sdp := <-webRtcStream.SDP:
				log.Println("Writing SDP")
				_, err := c.Writer.Write([]byte(sdp))
				if err != nil {
					log.Println("Writing SDP error", err)
					return
				}
			case <-webRtcStream.Ready:
				log.Println("webrtc ready")
				return
			}
		}

	}

	return

}

func (s *Stream) webServer() *sync.WaitGroup {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		router := gin.Default()
		router.LoadHTMLGlob("../examples/web/templates/*")
		router.GET("/", func(c *gin.Context) {
			firstUUID, allUUIDs := s.listStreams()
			log.Println("current open streams", allUUIDs)

			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"port":     ":9000",
				"suuid":    firstUUID,
				"suuidMap": allUUIDs,
				"version":  time.Now().String(),
			})
		})
		router.GET("/player/:suuid", func(c *gin.Context) {
			_, all := s.listStreams()
			sort.Strings(all)
			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"port":     ":9000",
				"suuid":    c.Param("suuid"),
				"suuidMap": all,
				"version":  time.Now().String(),
			})
		})
		router.POST("/recive", func(context *gin.Context) {
			fmt.Println("New Client connected")

			context.Header("Access-Control-Allow-Origin", "*")
			sdpData := context.PostForm("data")
			streamUUID := context.PostForm("streamUUID")

			var currentRtsp *network.RtspClient

			for _, rc := range s.RtspClients {
				if rc.UUID == streamUUID {
					currentRtsp = rc
					break
				}
			}

			if currentRtsp != nil {
				webRtcStream := network.NewWebRtcClient(currentRtsp)

				webRtcStream.OpenWebRtcStream(context, sdpData)
			}
		})
		router.GET("/codec/:uuid", func(c *gin.Context) {
			c.Header("Access-Control-Allow-Origin", "*")
			if s.RtspClients == nil {
				return
			}

			for _, rc := range s.RtspClients {
				if rc.UUID == c.Param("uuid") {
					b, err := json.Marshal(rc.Stream.Codecs)
					if err == nil {
						_, err = c.Writer.Write(b)
						if err == nil {
							return
						}
					}
					break
				}
			}

		})
		router.StaticFS("/static", http.Dir("../examples/web/static"))
		err := router.Run(":9000")
		if err != nil {
			log.Fatalln("Start HTTP Server error", err)
		}
	}()

	return wg
}

func (s *Stream) listStreams() (first string, streams []string) {
	for _, rc := range s.RtspClients {
		log.Printf("list rtsp stream %s", rc.UUID)
		if first == "" {
			first = rc.UUID
		}
		streams = append(streams, rc.UUID)
	}
	return
}
