package test

import (
	"encoding/json"
	"github.com/ReolinkCameraAPI/reolinkapigo/pkg/reolinkapi"
	"github.com/jarcoal/httpmock"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func registerMockZoomOperation() {
	httpmock.RegisterResponder("POST", "http://127.0.0.1/cgi-bin/api.cgi",
		func(req *http.Request) (*http.Response, error) {

			type ReqData struct {
				Cmd    string                     `json:"cmd"`
				Action int                        `json:"action"`
				Param  json.RawMessage `json:"param"`
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

			type zoomOps struct {
				Channel string `json:"channel"`
				OP      string `json:"op"`
			}

			var zoomOperation *zoomOps

			err = json.Unmarshal(reqData[0].Param, &zoomOperation)

			log.Printf("ptzctrl operation %v", zoomOperation)

			generalData := map[string]interface{}{
				"cmd":  "PtzCtrl",
				"code": 0,
				"value": map[string]interface{}{
					"rspCode": 200,
				},
			}

			return httpmock.NewJsonResponse(200, []interface{}{generalData})
		},
	)
}

func registerMockFocusOperation() {
	httpmock.RegisterResponder("POST", "http://127.0.0.1/cgi-bin/api.cgi",
		func(req *http.Request) (*http.Response, error) {

			type ReqData struct {
				Cmd    string                     `json:"cmd"`
				Action int                        `json:"action"`
				Param  json.RawMessage `json:"param"`
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

			type focusOps struct {
				Channel string `json:"channel"`
				OP      string `json:"op"`
			}

			var focusOperation *focusOps

			err = json.Unmarshal(reqData[0].Param, &focusOperation)

			log.Printf("ptzctrl operation %v", focusOperation)

			generalData := map[string]interface{}{
				"cmd":  "PtzCtrl",
				"code": 0,
				"value": map[string]interface{}{
					"rspCode": 200,
				},
			}

			return httpmock.NewJsonResponse(200, []interface{}{generalData})
		},
	)
}

func TestZoomMixin_StartZoomingIn(t *testing.T) {
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

	registerMockZoomOperation()

	ok, err := camera.StartZoomingIn()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("StartZoomIn %v", ok)
}

func TestZoomMixin_StartZoomingOut(t *testing.T) {
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

	registerMockZoomOperation()

	ok, err := camera.StartZoomingOut()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("StartZoomOut %v", ok)
}

func TestZoomMixin_StopZooming(t *testing.T) {
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

	registerMockZoomOperation()

	ok, err := camera.StopZooming()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("StopZooming %v", ok)
}

func TestZoomMixin_StartFocusingIn(t *testing.T) {
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

	registerMockFocusOperation()

	ok, err := camera.StartFocusingIn()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("FocusingIn %v", ok)
}

func TestZoomMixin_StartFocusingOut(t *testing.T) {
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

	registerMockFocusOperation()

	ok, err := camera.StartFocusingOut()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("FocusingOut %v", ok)
}

func TestZoomMixin_StopFocusing(t *testing.T) {
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

	registerMockFocusOperation()

	ok, err := camera.StopFocusing()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("StopFocusing %v", ok)
}
