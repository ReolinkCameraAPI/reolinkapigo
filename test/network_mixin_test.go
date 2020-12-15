package test

import (
	"encoding/json"
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/models"
	"github.com/ReolinkCameraAPI/reolink-go-api/pkg"
	"github.com/jarcoal/httpmock"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func registerMockSetNetworkPort() {
	httpmock.RegisterResponder("POST", "http://127.0.0.1/cgi-bin/api.cgi",
		func(req *http.Request) (*http.Response, error) {

			type ReqData struct {
				Cmd    string                     `json:"cmd"`
				Action int                        `json:"action"`
				Param  map[string]json.RawMessage `json:"param"`
			}

			// check the username and password
			var reqData []*ReqData

			data, err := ioutil.ReadAll(req.Body)

			if err != nil {
				return httpmock.NewStringResponse(500, err.Error()), nil
			}

			err = json.Unmarshal(data, &reqData)

			if err != nil {
				return httpmock.NewStringResponse(500, err.Error()), nil
			}

			var networkPort map[string]interface{}

			err = json.Unmarshal(reqData[0].Param["NetPort"], &networkPort)

			if err != nil {
				return httpmock.NewStringResponse(500, err.Error()), nil
			}

			log.Printf("received NetPort: %v", networkPort)

			generalData := map[string]interface{}{
				"cmd":  "SetNetPort",
				"code": 0,
				"value": map[string]interface{}{
					"rspCode": 200,
				},
			}

			return httpmock.NewJsonResponse(200, []interface{}{generalData})

		},
	)
}

func registerMockSetWifi() {
	httpmock.RegisterResponder("POST", "http://127.0.0.1/cgi-bin/api.cgi",
		func(req *http.Request) (*http.Response, error) {

			type ReqData struct {
				Cmd    string                     `json:"cmd"`
				Action int                        `json:"action"`
				Param  map[string]json.RawMessage `json:"param"`
			}

			// check the username and password
			var reqData []*ReqData

			data, err := ioutil.ReadAll(req.Body)

			if err != nil {
				return httpmock.NewStringResponse(500, err.Error()), nil
			}

			err = json.Unmarshal(data, &reqData)

			if err != nil {
				return httpmock.NewStringResponse(500, err.Error()), nil
			}

			var wifiSet map[string]interface{}

			err = json.Unmarshal(reqData[0].Param["Wifi"], &wifiSet)

			if err != nil {
				return httpmock.NewStringResponse(500, err.Error()), nil
			}

			log.Printf("received NetPort: %v", wifiSet)

			generalData := map[string]interface{}{
				"cmd":  "SetWifi",
				"code": 0,
				"value": map[string]interface{}{
					"rspCode": 200,
				},
			}

			return httpmock.NewJsonResponse(200, []interface{}{generalData})

		},
	)
}

func registerMockGetWifi() {
	httpmock.RegisterResponder("POST", "http://127.0.0.1/cgi-bin/api.cgi",
		func(req *http.Request) (*http.Response, error) {

			type ReqData struct {
				Cmd    string                     `json:"cmd"`
				Action int                        `json:"action"`
				Param  map[string]json.RawMessage `json:"param"`
			}

			// check the username and password
			var reqData []*ReqData

			data, err := ioutil.ReadAll(req.Body)

			if err != nil {
				return httpmock.NewStringResponse(500, err.Error()), nil
			}

			err = json.Unmarshal(data, &reqData)

			if err != nil {
				return httpmock.NewStringResponse(500, err.Error()), nil
			}

			wifiInfo := &models.Wifi{}

			generalData := map[string]interface{}{
				"cmd":  "GetWifi",
				"code": 0,
				"value": map[string]interface{}{
					"Wifi": wifiInfo,
				},
			}

			return httpmock.NewJsonResponse(200, []interface{}{generalData})

		},
	)
}

func registerMockGetNetworkGeneral() {
	httpmock.RegisterResponder("POST", "http://127.0.0.1/cgi-bin/api.cgi",
		func(req *http.Request) (*http.Response, error) {

			type ReqData struct {
				Cmd    string                     `json:"cmd"`
				Action int                        `json:"action"`
				Param  map[string]json.RawMessage `json:"param"`
			}

			// check the username and password
			var reqData []*ReqData

			data, err := ioutil.ReadAll(req.Body)

			if err != nil {
				return httpmock.NewStringResponse(500, err.Error()), nil
			}

			err = json.Unmarshal(data, &reqData)

			if err != nil {
				return httpmock.NewStringResponse(500, err.Error()), nil
			}

			networkGeneral := &models.NetworkGeneral{
				ActiveLink: "LAN",
				Dns: models.NetworkGeneralDns{
					Auto: 1,
					Dns1: "192.168.255.4",
					Dns2: "192.168.255.4",
				},
				Mac: "EC:71:DB:AA:59:CF",
				Static: models.NetworkGeneralStatic{
					Gateway: "192.168.255.1",
					Ip:      "192.168.255.58",
					Mask:    "255.255.255.0",
				},
				Type: "DHCP",
			}

			generalData := map[string]interface{}{
				"cmd":  "GetLocalLink",
				"code": 0,
				"value": map[string]interface{}{
					"LocalLink": networkGeneral,
				},
			}

			return httpmock.NewJsonResponse(200, []interface{}{generalData})

		},
	)
}

func registerMockScanWifi() {
	httpmock.RegisterResponder("POST", "http://127.0.0.1/cgi-bin/api.cgi",
		func(req *http.Request) (*http.Response, error) {

			type ReqData struct {
				Cmd    string                     `json:"cmd"`
				Action int                        `json:"action"`
				Param  map[string]json.RawMessage `json:"param"`
			}

			// check the username and password
			var reqData []*ReqData

			data, err := ioutil.ReadAll(req.Body)

			if err != nil {
				return httpmock.NewStringResponse(500, err.Error()), nil
			}

			err = json.Unmarshal(data, &reqData)

			if err != nil {
				return httpmock.NewStringResponse(500, err.Error()), nil
			}

			scanWifi := &models.ScanWifi{}

			generalData := map[string]interface{}{
				"cmd":  "ScanWifi",
				"code": 0,
				"value": map[string]interface{}{
					"ScanWifi": scanWifi,
				},
			}

			return httpmock.NewJsonResponse(200, []interface{}{generalData})

		},
	)
}

func registerMockGetNetworkDDNS() {
	httpmock.RegisterResponder("POST", "http://127.0.0.1/cgi-bin/api.cgi",
		func(req *http.Request) (*http.Response, error) {

			type ReqData struct {
				Cmd    string                     `json:"cmd"`
				Action int                        `json:"action"`
				Param  map[string]json.RawMessage `json:"param"`
			}

			// check the username and password
			var reqData []*ReqData

			data, err := ioutil.ReadAll(req.Body)

			if err != nil {
				return httpmock.NewStringResponse(500, err.Error()), nil
			}

			err = json.Unmarshal(data, &reqData)

			if err != nil {
				return httpmock.NewStringResponse(500, err.Error()), nil
			}

			networkDDNS := &models.NetworkDDNS{
				Domain:   "",
				Enable:   false,
				Password: "",
				Type:     "no-ip",
				Username: "",
			}

			generalData := map[string]interface{}{
				"cmd":  "GetDdns",
				"code": 0,
				"value": map[string]interface{}{
					"Ddns": networkDDNS,
				},
			}

			return httpmock.NewJsonResponse(200, []interface{}{generalData})

		},
	)
}

func registerMockGetNetworkNTP() {
	httpmock.RegisterResponder("POST", "http://127.0.0.1/cgi-bin/api.cgi",
		func(req *http.Request) (*http.Response, error) {

			type ReqData struct {
				Cmd    string                     `json:"cmd"`
				Action int                        `json:"action"`
				Param  map[string]json.RawMessage `json:"param"`
			}

			// check the username and password
			var reqData []*ReqData

			data, err := ioutil.ReadAll(req.Body)

			if err != nil {
				return httpmock.NewStringResponse(500, err.Error()), nil
			}

			err = json.Unmarshal(data, &reqData)

			if err != nil {
				return httpmock.NewStringResponse(500, err.Error()), nil
			}

			networkNtp := &models.NetworkNTP{
				Enable:   true,
				Interval: 1440,
				Port:     123,
				Server:   "ntp.moos.xyz",
			}

			generalData := map[string]interface{}{
				"cmd":  "GetNtp",
				"code": 0,
				"value": map[string]interface{}{
					"Ntp": networkNtp,
				},
			}

			return httpmock.NewJsonResponse(200, []interface{}{generalData})

		},
	)
}

func registerMockGetNetworkEmail() {
	httpmock.RegisterResponder("POST", "http://127.0.0.1/cgi-bin/api.cgi",
		func(req *http.Request) (*http.Response, error) {

			type ReqData struct {
				Cmd    string                     `json:"cmd"`
				Action int                        `json:"action"`
				Param  map[string]json.RawMessage `json:"param"`
			}

			// check the username and password
			var reqData []*ReqData

			data, err := ioutil.ReadAll(req.Body)

			if err != nil {
				return httpmock.NewStringResponse(500, err.Error()), nil
			}

			err = json.Unmarshal(data, &reqData)

			if err != nil {
				return httpmock.NewStringResponse(500, err.Error()), nil
			}

			networkEmail := &models.NetworkEmail{
				Username:   "",
				Password:   "",
				Addr1:      "",
				Addr2:      "",
				Addr3:      "",
				Attachment: "picture",
				Interval:   "5 Minute",
				Nickname:   "",
				Schedule: models.Schedule{
					Enable: true,
					Table:  "111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111",
				},
				SmtpPort:   465,
				SmtpServer: "smtp.gmail.com",
				SSL:        true,
			}

			generalData := map[string]interface{}{
				"cmd":  "GetEmail",
				"code": 0,
				"value": map[string]interface{}{
					"Email": networkEmail,
				},
			}

			return httpmock.NewJsonResponse(200, []interface{}{generalData})

		},
	)
}

func registerMockGetNetworkFTP() {
	httpmock.RegisterResponder("POST", "http://127.0.0.1/cgi-bin/api.cgi",
		func(req *http.Request) (*http.Response, error) {

			type ReqData struct {
				Cmd    string                     `json:"cmd"`
				Action int                        `json:"action"`
				Param  map[string]json.RawMessage `json:"param"`
			}

			// check the username and password
			var reqData []*ReqData

			data, err := ioutil.ReadAll(req.Body)

			if err != nil {
				return httpmock.NewStringResponse(500, err.Error()), nil
			}

			err = json.Unmarshal(data, &reqData)

			if err != nil {
				return httpmock.NewStringResponse(500, err.Error()), nil
			}

			networkFtp := &models.NetworkFTP{
				Username:  "",
				Password:  "",
				Anonymous: false,
				Interval:  30,
				MaxSize:   100,
				Mode:      0,
				Port:      21,
				RemoteDir: "",
				Schedule: models.Schedule{
					Enable: true,
					Table:  "111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111",
				},
				Server:     "",
				StreamType: 0,
			}

			generalData := map[string]interface{}{
				"cmd":  "GetFtp",
				"code": 0,
				"value": map[string]interface{}{
					"Ftp": networkFtp,
				},
			}

			return httpmock.NewJsonResponse(200, []interface{}{generalData})

		},
	)
}

func registerMockGetNetworkPush() {
	httpmock.RegisterResponder("POST", "http://127.0.0.1/cgi-bin/api.cgi",
		func(req *http.Request) (*http.Response, error) {

			type ReqData struct {
				Cmd    string                     `json:"cmd"`
				Action int                        `json:"action"`
				Param  map[string]json.RawMessage `json:"param"`
			}

			// check the username and password
			var reqData []*ReqData

			data, err := ioutil.ReadAll(req.Body)

			if err != nil {
				return httpmock.NewStringResponse(500, err.Error()), nil
			}

			err = json.Unmarshal(data, &reqData)

			if err != nil {
				return httpmock.NewStringResponse(500, err.Error()), nil
			}

			networkPush := &models.NetworkPush{
				Schedule: models.Schedule{
					Enable: true,
					Table:  "111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111",
				},
			}

			generalData := map[string]interface{}{
				"cmd":  "GetPush",
				"code": 0,
				"value": map[string]interface{}{
					"Push": networkPush,
				},
			}

			return httpmock.NewJsonResponse(200, []interface{}{generalData})

		},
	)
}

func TestNetworkMixin_SetNetworkPort(t *testing.T) {
	httpmock.Activate()

	defer httpmock.DeactivateAndReset()

	registerMockAuth()
	camera, err := pkg.NewCamera("foo", "bar", "127.0.0.1")

	if err != nil {
		t.Error(err)
	}

	if camera.RestHandler.Token == "12345" {
		t.Logf("login successful")
	}

	registerMockSetNetworkPort()

	ok, err := camera.API.SetNetworkPort()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("SetNetworkPort %v", ok)
}

func TestNetworkMixin_SetWifi(t *testing.T) {
	httpmock.Activate()

	defer httpmock.DeactivateAndReset()

	registerMockAuth()
	camera, err := pkg.NewCamera("foo", "bar", "127.0.0.1")

	if err != nil {
		t.Error(err)
	}

	if camera.RestHandler.Token == "12345" {
		t.Logf("login successful")
	}

	registerMockSetWifi()

	ok, err := camera.API.SetWifi("wifi_2G", "wow1234")(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("SetWifi %v", ok)

}

func TestNetworkMixin_GetWifi(t *testing.T) {
	httpmock.Activate()

	defer httpmock.DeactivateAndReset()

	registerMockAuth()
	camera, err := pkg.NewCamera("foo", "bar", "127.0.0.1")

	if err != nil {
		t.Error(err)
	}

	if camera.RestHandler.Token == "12345" {
		t.Logf("login successful")
	}

	registerMockGetWifi()

	wifiInfo, err := camera.API.GetWifi()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("GetWifi: %v", wifiInfo)
}

func TestNetworkMixin_ScanWifi(t *testing.T) {

	httpmock.Activate()

	defer httpmock.DeactivateAndReset()

	registerMockAuth()
	camera, err := pkg.NewCamera("foo", "bar", "127.0.0.1")

	if err != nil {
		t.Error(err)
	}

	if camera.RestHandler.Token == "12345" {
		t.Logf("login successful")
	}

	registerMockScanWifi()

	scanWifiInfo, err := camera.API.ScanWifi()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("ScanWifi %v", scanWifiInfo)
}

func TestNetworkMixin_GetNetworkGeneral(t *testing.T) {
	httpmock.Activate()

	defer httpmock.DeactivateAndReset()

	registerMockAuth()
	camera, err := pkg.NewCamera("foo", "bar", "127.0.0.1")

	if err != nil {
		t.Error(err)
	}

	if camera.RestHandler.Token == "12345" {
		t.Logf("login successful")
	}

	registerMockGetNetworkGeneral()

	networkGeneralInfo, err := camera.API.GetNetworkGeneral()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("GetNetworkGeneral %v", networkGeneralInfo)
}

func TestNetworkMixin_GetNetworkDDNS(t *testing.T) {
	httpmock.Activate()

	defer httpmock.DeactivateAndReset()

	registerMockAuth()

	camera, err := pkg.NewCamera("foo", "bar", "127.0.0.1")

	if err != nil {
		t.Error(err)
	}

	registerMockGetNetworkDDNS()

	networkDdns, err := camera.API.GetNetworkDDNS()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("GetNetworkDDNS %v", networkDdns)
}

func TestNetworkMixin_GetNetworkNTP(t *testing.T) {
	httpmock.Activate()

	defer httpmock.DeactivateAndReset()

	registerMockAuth()

	camera, err := pkg.NewCamera("foo", "bar", "127.0.0.1")

	if err != nil {
		t.Error(err)
	}

	registerMockGetNetworkNTP()

	networkNtp, err := camera.API.GetNetworkNTP()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("GetNetworkNTP %v", networkNtp)
}

func TestNetworkMixin_GetNetworkEmail(t *testing.T) {
	httpmock.Activate()

	defer httpmock.DeactivateAndReset()

	registerMockAuth()

	camera, err := pkg.NewCamera("foo", "bar", "127.0.0.1")

	if err != nil {
		t.Error(err)
	}

	registerMockGetNetworkEmail()

	networkNtp, err := camera.API.GetNetworkEmail()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("GetNetworkEmail %v", networkNtp)
}

func TestNetworkMixin_GetNetworkFTP(t *testing.T) {
	httpmock.Activate()

	defer httpmock.DeactivateAndReset()

	registerMockAuth()

	camera, err := pkg.NewCamera("foo", "bar", "127.0.0.1")

	if err != nil {
		t.Error(err)
	}

	registerMockGetNetworkFTP()

	networkFtp, err := camera.API.GetNetworkFTP()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("GetNetworkFTP %v", networkFtp)
}

func TestNetworkMixin_GetNetworkPush(t *testing.T) {
	httpmock.Activate()

	defer httpmock.DeactivateAndReset()

	registerMockAuth()

	camera, err := pkg.NewCamera("foo", "bar", "127.0.0.1")

	if err != nil {
		t.Error(err)
	}

	registerMockGetNetworkPush()

	networkPush, err := camera.API.GetNetworkPush()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("GetNetworkPush %v", networkPush)
}

func TestNetworkMixin_GetNetworkStatus(t *testing.T) {
	httpmock.Activate()

	defer httpmock.DeactivateAndReset()

	registerMockAuth()

	camera, err := pkg.NewCamera("foo", "bar", "127.0.0.1")

	if err != nil {
		t.Error(err)
	}

	registerMockGetNetworkGeneral()

	networkGeneral, err := camera.API.GetNetworkStatus()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("GetNetworkStatus %v", networkGeneral)
}
