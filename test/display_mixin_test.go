package test

import (
	"encoding/json"
	"github.com/ReolinkCameraAPI/reolinkapigo/internal/pkg/enum"
	"github.com/ReolinkCameraAPI/reolinkapigo/internal/pkg/models"
	"github.com/ReolinkCameraAPI/reolinkapigo/pkg"
	"github.com/jarcoal/httpmock"
	"io/ioutil"
	"net/http"
	"testing"
)

func registerMockGetOsd() {
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

			osdInfo := &models.Osd{
				BgColor: false,
				Channel: 0,
				OsdChannel: models.OsdChannel{
					Enable: false,
					Name:   "FarRight",
					Pos:    enum.LOWER_RIGHT.Value(),
				},
				OsdTime: models.OsdTime{
					Enable: false,
					Pos:    enum.LOWER_LEFT.Value(),
				},
			}

			generalData := map[string]interface{}{
				"cmd":  "GetOsd",
				"code": 0,
				"value": map[string]interface{}{
					"Osd": osdInfo,
				},
			}

			return httpmock.NewJsonResponse(200, []interface{}{generalData})

		},
	)
}

func registerMockGetMask() {
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

			maskInfo := &models.MaskData{
				Area: []models.MaskArea{
					{
						Block: models.MaskAreaBlock{
							Height: 128,
							Width:  250,
							X:      100,
							Y:      50,
						},
						Screen: models.MaskAreaScreen{
							Height: 128,
							Width:  250,
						},
					},
				},
				Channel: 0,
				Enable:  false,
			}

			generalData := map[string]interface{}{
				"cmd":  "GetMask",
				"code": 0,
				"value": map[string]interface{}{
					"Mask": maskInfo,
				},
			}

			return httpmock.NewJsonResponse(200, []interface{}{generalData})

		},
	)
}

func registerMockSetOsd() {
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

			var osdData *models.Osd

			err = json.Unmarshal(reqData[0].Param["Osd"], &osdData)

			if err != nil {
				return httpmock.NewStringResponse(500, err.Error()), nil
			}

			generalData := map[string]interface{}{
				"cmd":  "SetOsd",
				"code": 0,
				"value": map[string]interface{}{
					"rspCode": 200,
				},
			}

			return httpmock.NewJsonResponse(200, []interface{}{generalData})

		},
	)
}

func TestDisplayMixin_GetOSD(t *testing.T) {
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

	registerMockGetOsd()

	osdInfo, err := camera.API.GetOSD()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("%v", osdInfo)
}

func TestDisplayMixin_GetMask(t *testing.T) {
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

	registerMockGetMask()

	maskInfo, err := camera.API.GetMask()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("%v", maskInfo)
}

func TestDisplayMixin_SetOSD(t *testing.T) {
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

	registerMockSetOsd()

	ok, err := camera.API.SetOSD()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("SetOSD %v", ok)


}
