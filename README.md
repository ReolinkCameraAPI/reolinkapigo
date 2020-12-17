<h1 align="center">Reolink Go Api Client</h1>

<p align="center">
    	<img alt="Reolink Approval" src="https://img.shields.io/badge/reolink-approved-blue?style=flat-square">
	<img alt="GitHub" src="https://img.shields.io/github/license/ReolinkCameraAPI/reolink-go-api?style=flat-square">
	<img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/ReolinkCameraAPI/reolink-go-api?style=flat-square">
	<img alt="GitHub tag (latest SemVer)" src="https://img.shields.io/github/v/tag/ReolinkCameraApi/reolink-go-api?style=flat-square">
   	<img alt="Discord" src="https://img.shields.io/discord/773257004911034389?style=flat-square">
</p>

---

A Reolink Camera client written in Go. This repository's 
purpose **(with Reolink's full support)** is to deliver a 
complete API for the Reolink Cameras,
although they have a basic API document - it does not satisfy the 
need for extensive camera communication.

Check out our documentation for more information on how to use the software at [https://reolink.oleaintueri.com](https://oleaintueri.com)

Other Supported Languages:

- Python: [reolink-python-api](https://github.com/ReolinkCameraApi/reolink-python-api)

The reolink-go-api project is the go alternative to the reolink-python-api project. It provides the same functionality,
but just in pure Go.

### WARNING...This is an untested repository and is in heavy development

### Join us on Discord

    https://discord.gg/8z3fdAmZJP

### Sponsorship

<a href="https://oleaintueri.com"><img src="https://oleaintueri.com/images/oliv.svg" width="60px"/><img width="200px" style="padding-bottom: 10px" src="https://oleaintueri.com/images/oleaintueri.svg"/></a>

[Oleaintueri](https://oleaintueri.com) is sponsoring the development and maintenance of these projects within their organisation.


---

## Get started

### Installation

    go get -u github.com/ReolinkCameraApi/reolink-go-api

### Usage

Check the `examples/` directory for example code and implementations.

Implement a "Camera" object by passing it an IP address, Username and Password. By instantiating the object, it will try
retrieve a login token from the Reolink Camera. This token is necessary to interact with the Camera using other
commands.

    import "github.com/ReolinkCameraApi/reolinkapigo"

    // This can throw an error due to the API trying to authorise with the camera
    // to retrieve the necessary token for future requests.
	camera, err := pkg.NewCamera("foo", "bar", "192.168.1.100")

    // now call any of the supported api's by passing it it's resthandler
    ok, err := camera.API.FormatHdd(0)(camera.RestHandler)

Dependencies needed to make this work:

- [Deepch's VDK](github.com/deepch/vdk)
- [Pions's WebRTC](github.com/pion/webrtc/v2)

Dependencies needed for testing:
- [Gin](github.com/gin-gonic/gin)
- [HttpMock](github.com/jarcoal/httpmock)


## Contributors

---
### Styling and Standards

Golang project structure based off of https://github.com/golang-standards/project-layout

Writing functions/structs with a lot of parameters

    function HasManyParameters(
            param1 string,
            param2 int,
            param3 bool,
            ...
            param100 *Foo,
    ) {
       // Write your code here
    }

All variables are camelCase

    var someVariable1 string

Package names are all lowercase and if two or more words, camelCase

    package foo
    package fooBar

Go files are lowercase and if two or more words, snake_case

    foo.go
    foo_bar.go

### How can I become a contributor?

#### Step 1

Get the Restful API calls by looking through the HTTP Requests made in the camera's web UI. I use Google Chrome developer mode (ctr + shift + i) -> Network.

#### Step 2

Fork the repository and make your changes.

#### Step 3

Make a pull request.

### Test without a camera

All the tests implement a MockApi. The only test that could be a bit tricky is the RTSP client test.

To test this locally on your machine you could use the [rtsp-simple-server](https://github.com/aler9/rtsp-simple-server)


Setting it up is quite easy, however streaming the video feed needs some extra thought, especially if you are new to
ffmpeg.

Steps:
- Get rtsp-simple-server (download binary etc.) and Start server.
- Find a video file and push content to server
- Run the rtsp test
    
In your terminal:

    //receives the stream and passes it along to clients
    
    ./rtsp-simple-server
    
    // this will start and encode the stream on the fly
    
    ffmpeg -re -stream_loop -1 \ 
    -i vidfile.mkv \
    -c:a aac -b:a 64k -c:v libx264 -preset ultrafast -b:v 500k -f hls -hls_time 1 -hls_list_size 3 -hls_flags delete_segments -hls_allow_cache 0 \
    -f rtsp rtsp://localhost:8554/mystream


### API Requests Implementation Plan:

Stream:
- [X] RTSP
- [X] WebRTC

GET:

- [X] Login
- [ ] Logout
- [ ] Display -> OSD
- [ ] Recording -> Encode (Clear and Fluent Stream)
- [ ] Recording -> Advance (Scheduling)
- [ ] Network -> General
- [ ] Network -> Advanced
- [ ] Network -> DDNS
- [ ] Network -> NTP
- [ ] Network -> E-mail
- [ ] Network -> FTP
- [ ] Network -> Push
- [ ] Network -> WIFI
- [ ] Alarm -> Motion
- [ ] System -> General
- [ ] System -> DST
- [ ] System -> Information
- [ ] System -> Maintenance
- [ ] System -> Performance
- [ ] System -> Reboot
- [ ] User -> Online User
- [ ] User -> Add User
- [ ] User -> Manage User
- [ ] Device -> HDD/SD Card
- [ ] Zoom
- [ ] Focus
- [ ] Image (Brightness, Contrast, Saturation, Hue, Sharp, Mirror, Rotate)
- [ ] Advanced Image (Anti-flicker, Exposure, White Balance, DayNight, Backlight, LED light, 3D-NR)
- [ ] Image Data -> "Snap" Frame from Video Stream

SET:

- [ ] Display -> OSD
- [ ] Recording -> Encode (Clear and Fluent Stream)
- [ ] Recording -> Advance (Scheduling)
- [ ] Network -> General
- [ ] Network -> Advanced
- [ ] Network -> DDNS
- [ ] Network -> NTP
- [ ] Network -> E-mail
- [ ] Network -> FTP
- [ ] Network -> Push
- [ ] Network -> WIFI
- [ ] Alarm -> Motion
- [ ] System -> General
- [ ] System -> DST
- [ ] System -> Reboot
- [ ] User -> Online User
- [ ] User -> Add User
- [ ] User -> Manage User
- [ ] Device -> HDD/SD Card (Format)
- [ ] PTZ
- [ ] Zoom
- [ ] Focus
- [ ] Image (Brightness, Contrast, Saturation, Hue, Sharp, Mirror, Rotate)
- [ ] Advanced Image (Anti-flicker, Exposure, White Balance, DayNight, Backlight, LED light, 3D-NR)

### Supported Camera's

Any Reolink camera that has a web UI should work. The other's requiring special Reolink clients
do not work and is not supported here.

- RLC-411WS
- RLC-423
- RLC-420-5MP
- RLC-410-5MP
- RLC-520

### Integrated source code from:

- [RTSPtoWebRTC by Deepch](https://github.com/deepch/RTSPtoWebRTC)
