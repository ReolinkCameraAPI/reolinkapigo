package test

import (
	"encoding/json"
	"github.com/ReolinkCameraAPI/reolinkapigo/internal/pkg/api"
	"github.com/ReolinkCameraAPI/reolinkapigo/internal/pkg/enum"
	"github.com/ReolinkCameraAPI/reolinkapigo/pkg/reolinkapi"
	"github.com/jarcoal/httpmock"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func registerMockSetAdvancedImage() {
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

			var advImgSet map[string]interface{}

			err = json.Unmarshal(reqData[0].Param["Isp"], &advImgSet)

			if err != nil {
				return httpmock.NewStringResponse(500, err.Error()), nil
			}

			log.Printf("received Isp: %v", advImgSet)

			generalData := map[string]interface{}{
				"cmd":  "SetIsp",
				"code": 0,
				"value": map[string]interface{}{
					"rspCode": 200,
				},
			}

			return httpmock.NewJsonResponse(200, []interface{}{generalData})

		},
	)
}

func registerMockSetImageSettings() {
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

			var imgSet map[string]interface{}

			err = json.Unmarshal(reqData[0].Param["Image"], &imgSet)

			if err != nil {
				return httpmock.NewStringResponse(500, err.Error()), nil
			}

			log.Printf("received Image: %v", imgSet)

			generalData := map[string]interface{}{
				"cmd":  "SetImage",
				"code": 0,
				"value": map[string]interface{}{
					"rspCode": 200,
				},
			}

			return httpmock.NewJsonResponse(200, []interface{}{generalData})

		},
	)
}

func TestImageMixin_SetAdvanceImageSettings(t *testing.T) {
	httpmock.Activate()

	defer httpmock.DeactivateAndReset()

	registerMockAuth()

	camera, err := reolinkapi.NewCamera("127.0.0.1", reolinkapi.WithUsername("foo"), reolinkapi.WithPassword("bar"))

	if err != nil {
		t.Error(err)
	}

	if camera.GetToken() == "12345" {
		t.Logf("login successful")
	}

	registerMockSetAdvancedImage()

	ok, err := camera.SetAdvanceImageSettings(
		api.ImageAdvancedOptionDayNight(enum.DAY_NIGHT_AUTO),
		api.ImageAdvancedOptionBacklight(enum.DYNAMIC_RANGE_CONTROL),
		api.ImageAdvancedOptionBlc(1),
		)(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("SetAdvanceImageSettings %v", ok)

}

func TestImageMixin_SetImageSettings(t *testing.T) {
	httpmock.Activate()

	defer httpmock.DeactivateAndReset()

	registerMockAuth()

	camera, err := reolinkapi.NewCamera("127.0.0.1", reolinkapi.WithUsername("foo"), reolinkapi.WithPassword("bar"))

	if err != nil {
		t.Error(err)
	}

	if camera.GetToken() == "12345" {
		t.Logf("login successful")
	}

	registerMockSetImageSettings()

	ok, err := camera.SetImageSettings()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("SetImageSettings %v", ok)
}
