package test

import (
	"encoding/json"
	"github.com/ReolinkCameraAPI/reolinkapigo/internal/pkg/models"
	"github.com/ReolinkCameraAPI/reolinkapigo/pkg"
	"github.com/jarcoal/httpmock"
	"io/ioutil"
	"net/http"
	"testing"
)

func registerMockGetGeneralSystem() {
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

			if reqData[0].Cmd == "GetTime" {
				systemDst := &models.DstInformation{
					Enable:       true,
					EndHour:      1,
					EndMin:       0,
					EndMon:       11,
					EndSec:       0,
					EndWeek:      1,
					EndWeekday:   0,
					Offset:       1,
					StartHour:    2,
					StartMin:     0,
					StartMon:     3,
					StartSec:     0,
					StartWeek:    1,
					StartWeekday: 0,
				}

				systemTime := &models.TimeInformation{
					Day:      1,
					Hour:     15,
					HourFmt:  false,
					Min:      33,
					Mon:      12,
					Sec:      58,
					TimeFmt:  "DD/MM/YYYY",
					TimeZone: 21600,
					Year:     2020,
				}

				generalData := map[string]interface{}{
					"cmd":  "GetTime",
					"code": 0,
					"value": map[string]interface{}{
						"Dst":  systemDst,
						"Time": systemTime,
					},
				}

				return httpmock.NewJsonResponse(200, []interface{}{generalData})
			}

			if reqData[0].Cmd == "GetNorm" {

				generalData := map[string]interface{}{
					"cmd":  "GetNorm",
					"code": 0,
					"value": map[string]interface{}{
						"norm": "NTSC",
					},
				}

				return httpmock.NewJsonResponse(200, []interface{}{generalData})
			}

			return httpmock.NewStringResponse(500, "Operation Unknown"), nil

		},
	)
}

func registerMockGetPerformance() {
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

			systemPerformance := &models.DevicePerformanceInformation{
				CodecRate:     2154,
				CpuUsed:       14,
				NetThroughput: 0,
			}

			generalData := map[string]interface{}{
				"cmd":  "GetPerformance",
				"code": 0,
				"value": map[string]interface{}{
					"Performance": systemPerformance,
				},
			}

			return httpmock.NewJsonResponse(200, []interface{}{generalData})
		},
	)
}

func registerMockDeviceInformation() {
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

			deviceInformation := &models.DeviceInformation{
				B485:            0,
				IoInputNumber:   0,
				IoOutputNumber:  0,
				AudioNumber:     0,
				BuildDay:        "build 18081408",
				ConfigVersion:   "v2.0.0.0",
				ChannelNumber:   1,
				Detail:          "IPC_3816M100000000100000",
				DiskNumber:      1,
				FirmwareVersion: "v2.0.0.1389_18081408",
				HardwareVersion: "IPC_3816M",
				Model:           "RLC-411WS",
				Name:            "Camera1_withpersonality",
				Serial:          "00000000000000",
				Type:            "IPC",
				Wifi:            true,
			}

			generalData := map[string]interface{}{
				"cmd":  "GetDevInfo",
				"code": 0,
				"value": map[string]interface{}{
					"DevInfo": deviceInformation,
				},
			}

			return httpmock.NewJsonResponse(200, []interface{}{generalData})
		},
	)
}

func registerMockRebootCamera() {
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

			generalData := map[string]interface{}{
				"cmd":  "Reboot",
				"code": 0,
				"value": map[string]interface{}{
					"rspCode": 200,
				},
			}

			return httpmock.NewJsonResponse(200, []interface{}{generalData})
		},
	)
}

func TestRecordingMixin_GetGeneralSystem(t *testing.T) {
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

	registerMockGetGeneralSystem()

	systemInfo, err := camera.API.GetGeneralSystem()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	data, _ := json.Marshal(systemInfo)

	t.Logf("GetGeneralSystem %v", string(data))
}

func TestRecordingMixin_GetPerformance(t *testing.T) {
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

	registerMockGetPerformance()

	performanceInfo, err := camera.API.GetPerformance()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	data, err := json.Marshal(performanceInfo)

	t.Logf("GetPerformance %s", string(data))
}

func TestRecordingMixin_GetDeviceInformation(t *testing.T) {
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

	registerMockDeviceInformation()

	deviceInfo, err := camera.API.GetDeviceInformation()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	data, err := json.Marshal(deviceInfo)

	t.Logf("GetDeviceInformation %s", string(data))
}

func TestRecordingMixin_RebootCamera(t *testing.T) {
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

	registerMockRebootCamera()

	ok, err := camera.API.RebootCamera()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("RebootCamera %v", ok)

}

func TestRecordingMixin_GetDstInformation(t *testing.T) {
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

	registerMockGetGeneralSystem()

	dstInfo, timeInfo, err := camera.API.GetDstInformation()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	dataDst, _ := json.Marshal(dstInfo)

	dataTime, _ := json.Marshal(timeInfo)

	t.Logf("GetDstInformation, Dst %s", string(dataDst))

	t.Logf("GetDstInformation, Time %s", string(dataTime))
}
