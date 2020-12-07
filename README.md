## Reolink Go Api Client

A Reolink Camera client written in Go.

Other Supported Languages:
 - Python: [reolink-python-api](https://github.com/ReolinkCameraApi/reolink-python-api)
 
The reolink-go-api project is the go alternative to the reolink-python-api project.
It provides the same functionality, but just in pure Go.

### WARNING...This is an untested repository and is in heavy development

### Join us on Discord

    https://discord.gg/8z3fdAmZJP

### Purpose

This repository's purpose is to deliver a complete API for the Reolink Camera's, ( TESTED on RLC-411WS )


### But Reolink gives an API in their documentation

Not really. They only deliver a really basic API to retrieve Image data and Video data.

### How?

You can get the Restful API calls by looking through the HTTP Requests made the camera web console. I use Google Chrome developer mode (ctr + shift + i) -> Network.

### Get started

Implement a "Camera" object by passing it an IP address, Username and Password. 
By instantiating the object, it will try retrieve a login token from the Reolink Camera.
This token is necessary to interact with the Camera using other commands.

Dependencies needed to make this work:

 - [GoCV](https://gocv.io)

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

    

### API Requests Implementation Plan:

GET:
- [ ] Login
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