package test

import (
	"encoding/json"
	"github.com/ReolinkCameraAPI/reolinkapi/internal/pkg/api"
	"github.com/ReolinkCameraAPI/reolinkapi/pkg"
	"github.com/jarcoal/httpmock"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func registerMockGoToPreset() {
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

			ptzPreset := reqData[0].Param

			log.Printf("received PtzPreset: %v", ptzPreset)

			generalData := map[string]interface{}{
				"cmd":  "SetPtzPreset",
				"code": 0,
				"value": map[string]interface{}{
					"rspCode": 200,
				},
			}

			return httpmock.NewJsonResponse(200, []interface{}{generalData})

		},
	)
}

func registerMockAddPreset() {
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

			ptzPreset := reqData[0].Param

			log.Printf("received PtzPreset: %v", ptzPreset)

			generalData := map[string]interface{}{
				"cmd":  "SetPtzPreset",
				"code": 0,
				"value": map[string]interface{}{
					"rspCode": 200,
				},
			}

			return httpmock.NewJsonResponse(200, []interface{}{generalData})

		},
	)
}

func registerMockRemovePreset() {
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

			ptzPreset := reqData[0].Param

			log.Printf("received PtzPreset: %v", ptzPreset)

			generalData := map[string]interface{}{
				"cmd":  "SetPtzPreset",
				"code": 0,
				"value": map[string]interface{}{
					"rspCode": 200,
				},
			}

			return httpmock.NewJsonResponse(200, []interface{}{generalData})

		},
	)
}

func registerMockPtzOperation() {
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

			ptzPreset := reqData[0].Param

			log.Printf("received Ptz Operation: %v", ptzPreset)

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

func TestPtzMixin_GoToPreset(t *testing.T) {
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

	registerMockGoToPreset()

	ok, err := camera.API.GoToPreset(api.PtzOptionOpsIndex(nil), api.PtzOptionOpsSpeed(30))(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("SetNetworkPort %v", ok)
}

func TestPtzMixin_AddPreset(t *testing.T) {
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

	registerMockAddPreset()

	ok, err := camera.API.AddPreset(api.PtzOptionsPresetName("NewPtzPreset"))(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("AddPreset %v", ok)
}

func TestPtzMixin_RemovePreset(t *testing.T) {
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

	registerMockRemovePreset()

	ok, err := camera.API.RemovePreset(api.PtzOptionsPresetName("NewPtzPreset"))(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("AddPreset %v", ok)

}

func TestPtzMixin_MoveRight(t *testing.T) {
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

	registerMockPtzOperation()

	ok, err := camera.API.MoveRight()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("MoveRight %v", ok)
}

func TestPtzMixin_MoveRightUp(t *testing.T) {
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

	registerMockPtzOperation()

	ok, err := camera.API.MoveRightUp()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("MoveRightUp %v", ok)
}

func TestPtzMixin_MoveRightDown(t *testing.T) {
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

	registerMockPtzOperation()

	ok, err := camera.API.MoveRightDown()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("MoveRightDown %v", ok)
}

func TestPtzMixin_MoveLeft(t *testing.T) {
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

	registerMockPtzOperation()

	ok, err := camera.API.MoveLeft()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("MoveLeft %v", ok)
}

func TestPtzMixin_MoveLeftUp(t *testing.T) {
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

	registerMockPtzOperation()

	ok, err := camera.API.MoveLeftUp()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("MoveLeftUp %v", ok)
}

func TestPtzMixin_MoveLeftDown(t *testing.T) {
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

	registerMockPtzOperation()

	ok, err := camera.API.MoveLeftDown()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("MoveLeftDown %v", ok)
}

func TestPtzMixin_MoveUp(t *testing.T) {
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

	registerMockPtzOperation()

	ok, err := camera.API.MoveUp()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("MoveUp %v", ok)
}

func TestPtzMixin_MoveDown(t *testing.T) {
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

	registerMockPtzOperation()

	ok, err := camera.API.MoveDown()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("MoveDown %v", ok)
}

func TestPtzMixin_StopPtz(t *testing.T) {
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

	registerMockPtzOperation()

	ok, err := camera.API.StopPtz()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("Stop %v", ok)
}

func TestPtzMixin_AutoMovement(t *testing.T) {
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

	registerMockPtzOperation()

	ok, err := camera.API.AutoMovement()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("AutoMovement %v", ok)
}
