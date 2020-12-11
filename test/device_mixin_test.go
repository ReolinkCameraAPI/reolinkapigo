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

func registerMockHddInfo() {
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

			deviceHddInfo := &models.HddInfo{
				Capacity: 15181,
				Format:   1,
				ID:       0,
				Mount:    1,
				Size:     15181,
			}

			generalData := map[string]interface{}{
				"cmd":  "GetHddInfo",
				"code": 0,
				"value": map[string]interface{}{
					"HddInfo": deviceHddInfo,
				},
			}

			return httpmock.NewJsonResponse(200, []interface{}{generalData})

		},
	)
}

func registerMockFormatHdd() {
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

			formatInfo := map[string]interface{}{
				"rspCode": 200,
			}

			generalData := map[string]interface{}{
				"cmd":   "Format",
				"code":  0,
				"value": formatInfo,
			}

			return httpmock.NewJsonResponse(200, []interface{}{generalData})

		},
	)
}

func TestDeviceMixin_GetHddInfo(t *testing.T) {

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

	registerMockHddInfo()

	hddInfo, err := camera.API.GetHddInfo()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	log.Printf("%v", hddInfo)
}

func TestFormatMixin_FormatHdd(t *testing.T) {
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

	registerMockFormatHdd()

	formatInfo, err := camera.API.FormatHdd(0)(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	log.Printf("%v", formatInfo)
}
