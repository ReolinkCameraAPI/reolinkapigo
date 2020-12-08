package api

import (
	"encoding/json"
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/models"
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/network"
)

type NetworkMixin struct {
}

type networkPorts struct {
	http  int
	https int
	media int
	onvif int
	rtmp  int
	rtsp  int
}

// Set the camera network ports using the NetworkPortOption<prop> functions
// Defaults are automatically for the excluded networkPortOptions
// http: 80
// https: 443
// media: 9000
// onvif: 8000
// rtmp: 1935
// rtsp: 554
func (nm *NetworkMixin) SetNetworkPort(networkPortOptions ...func(nm *networkPorts) error) func(handler *network.RestHandler) (bool,
	error) {

	// Defaults
	networkPorts := &networkPorts{
		http:  80,
		https: 443,
		media: 9000,
		onvif: 8000,
		rtmp:  1935,
		rtsp:  554,
	}

	for _, op := range networkPortOptions {
		err := op(networkPorts)
		if err != nil {

		}
	}

	return func(handler *network.RestHandler) (bool, error) {
		payload := map[string]interface{}{
			"cmd":    "SetNetPort",
			"action": 0,
			"params": map[string]interface{}{
				"NetPort": map[string]interface{}{
					"httpPort":  networkPorts.http,
					"httpsPort": networkPorts.https,
					"mediaPort": networkPorts.media,
					"onvifPort": networkPorts.onvif,
					"rtmpPort":  networkPorts.rtmp,
					"rtspPort":  networkPorts.rtsp,
				},
			},
		}

		_, err := handler.Request("POST", payload, true)

		if err != nil {
			return false, err
		}

		// TODO: get the correct type back from the response.
		return true, nil
	}
}

// Set the camera's wifi settings
func (nm *NetworkMixin) SetWifi(ssid string, password string) func(handler *network.RestHandler) (bool, error) {
	return func(handler *network.RestHandler) (bool, error) {
		payload := map[string]interface{}{
			"cmd":    "SetWifi",
			"action": 0,
			"params": map[string]interface{}{
				"Wifi": map[string]interface{}{
					"ssid":     ssid,
					"password": password,
				},
			},
		}

		_, err := handler.Request("POST", payload, true)

		if err != nil {
			return false, err
		}

		// TODO: get the correct type back from the response.
		return true, nil
	}
}

// Get the current camera's wifi settings
func (nm *NetworkMixin) GetWifi() func(handler *network.RestHandler) (*models.Wifi, error) {
	return func(handler *network.RestHandler) (*models.Wifi, error) {
		payload := map[string]interface{}{
			"cmd":    "GetWifi",
			"action": 1,
			"params": map[string]interface{}{},
		}

		result, err := handler.Request("POST", payload, true)

		if err != nil {
			return nil, err
		}

		var wifi *models.Wifi

		err = json.Unmarshal(result.Value["Wifi"], &wifi)

		if err != nil {
			return nil, err
		}

		return wifi, nil
	}
}

// Scan the current camera's wifi network
func (nm *NetworkMixin) ScanWifi() func(handler *network.RestHandler) (*models.ScanWifi, error) {
	return func(handler *network.RestHandler) (*models.ScanWifi, error) {
		payload := map[string]interface{}{
			"cmd":    "ScanWifi",
			"action": 1,
			"params": map[string]interface{}{},
		}

		result, err := handler.Request("POST", payload, true)

		if err != nil {
			return nil, err
		}

		var scanWifi *models.ScanWifi

		err = json.Unmarshal(result.Value["ScanWifi"], &scanWifi)

		if err != nil {
			return nil, err
		}

		return scanWifi, nil
	}
}

// Get the camera's general network information
func (nm *NetworkMixin) GetNetworkGeneral() func(handler *network.RestHandler) (*models.NetworkGeneral, error) {
	return func(handler *network.RestHandler) (*models.NetworkGeneral, error) {
		payload := map[string]interface{}{
			"cmd":    "GetLocalLink",
			"action": 0,
			"params": map[string]interface{}{},
		}

		resp, err := handler.Request("POST", payload, true)

		if err != nil {
			return nil, err
		}

		var networkGeneral *models.NetworkGeneral

		err = json.Unmarshal(resp.Value["LocalLink"], &networkGeneral)

		if err != nil {
			return nil, err
		}

		return networkGeneral, nil
	}
}

// Get the camera's network DDNS information
func (nm *NetworkMixin) GetNetworkDDNS() func(handler *network.RestHandler) (*models.NetworkDDNS, error) {
	return func(handler *network.RestHandler) (*models.NetworkDDNS, error) {
		payload := map[string]interface{}{
			"cmd":    "GetDdns",
			"action": 0,
			"params": map[string]interface{}{},
		}

		resp, err := handler.Request("POST", payload, true)

		if err != nil {
			return nil, err
		}

		var networkDdns *models.NetworkDDNS

		err = json.Unmarshal(resp.Value["Ddns"], &networkDdns)

		if err != nil {
			return nil, err
		}

		return networkDdns, nil
	}
}

// Get the camera's network NTP information
func (nm *NetworkMixin) GetNetworkNTP() func(handler *network.RestHandler) (*models.NetworkNTP, error) {
	return func(handler *network.RestHandler) (*models.NetworkNTP, error) {
		payload := map[string]interface{}{
			"cmd":    "GetNtp",
			"action": 0,
			"params": map[string]interface{}{},
		}

		resp, err := handler.Request("POST", payload, true)

		if err != nil {
			return nil, err
		}

		var networkNtp *models.NetworkNTP

		err = json.Unmarshal(resp.Value["Ntp"], &networkNtp)

		if err != nil {
			return nil, err
		}

		return networkNtp, nil
	}
}

// Get the camera's network Email information
func (nm *NetworkMixin) GetNetworkEmail() func(handler *network.RestHandler) (*models.NetworkEmail, error) {
	return func(handler *network.RestHandler) (*models.NetworkEmail, error) {
		payload := map[string]interface{}{
			"cmd":    "GetEmail",
			"action": 0,
			"params": map[string]interface{}{},
		}

		resp, err := handler.Request("POST", payload, true)

		if err != nil {
			return nil, err
		}

		var networkEmail *models.NetworkEmail

		err = json.Unmarshal(resp.Value["Email"], &networkEmail)

		if err != nil {
			return nil, err
		}

		return networkEmail, nil
	}
}

// Get the camera's network FTP information
func (nm *NetworkMixin) GetNetworkFTP() func(handler *network.RestHandler) (*models.NetworkFTP, error) {
	return func(handler *network.RestHandler) (*models.NetworkFTP, error) {
		payload := map[string]interface{}{
			"cmd":    "GetFtp",
			"action": 0,
			"params": map[string]interface{}{},
		}

		resp, err := handler.Request("POST", payload, true)

		if err != nil {
			return nil, err
		}

		var networkFtp *models.NetworkFTP

		err = json.Unmarshal(resp.Value["Ftp"], &networkFtp)

		if err != nil {
			return nil, err
		}

		return networkFtp, nil
	}
}

// Get the camera's network Push information
func (nm *NetworkMixin) GetNetworkPush() func(handler *network.RestHandler) (*models.NetworkPush, error) {
	return func(handler *network.RestHandler) (*models.NetworkPush, error) {
		payload := map[string]interface{}{
			"cmd":    "GetPush",
			"action": 0,
			"params": map[string]interface{}{},
		}

		resp, err := handler.Request("POST", payload, true)

		if err != nil {
			return nil, err
		}

		var networkPush *models.NetworkPush

		err = json.Unmarshal(resp.Value["Push"], &networkPush)

		if err != nil {
			return nil, err
		}

		return networkPush, nil
	}
}

// Get the camera's network Status information is just a wrapper for networkGeneral
// TODO: revise this, exactly copied from the reolink-python-api project.
func (nm *NetworkMixin) GetNetworkStatus() func(handler *network.RestHandler) (*models.NetworkGeneral, error) {
	return func(handler *network.RestHandler) (*models.NetworkGeneral, error) {
		return nm.GetNetworkGeneral()(handler)
	}
}

// An option for SetNetworkPort to set the httpPort
// Default value of httpPort is 80
func NetworkPortOptionHttp(httpPort int) func(nm *networkPorts) error {
	return func(nm *networkPorts) error {
		nm.http = httpPort
		return nil
	}
}

// An option for SetNetworkPort to set the httpsPort
// Default value of httpsPort is 443
func NetworkPortOptionHttps(https int) func(nm *networkPorts) error {
	return func(nm *networkPorts) error {
		nm.https = https
		return nil
	}
}

// An option for SetNetworkPort to set the mediaPort
// Default value of mediaPort is 9000
func NetworkPortOptionMedia(media int) func(nm *networkPorts) error {
	return func(nm *networkPorts) error {
		nm.media = media
		return nil
	}
}

// An option for SetNetworkPort to set the onvifPort
// Default value of onvifPort is 8000
func NetworkPortOptionOnvif(onvif int) func(nm *networkPorts) error {
	return func(nm *networkPorts) error {
		nm.onvif = onvif
		return nil
	}
}

// An option for SetNetworkPort to set the rtmpPort
// Default value of rtmpPort is 1935
func NetworkPortOptionRtmp(rtmp int) func(nm *networkPorts) error {
	return func(nm *networkPorts) error {
		nm.rtmp = rtmp
		return nil
	}
}

// An option for SetNetworkPort to set the rtspPort
// Default value of rtspPort is 554
func NetworkPortOptionRtsp(rtsp int) func(nm *networkPorts) error {
	return func(nm *networkPorts) error {
		nm.rtsp = rtsp
		return nil
	}
}
