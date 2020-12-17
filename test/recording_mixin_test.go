package test

import (
	"encoding/json"
	"github.com/ReolinkCameraAPI/reolinkapi/internal/pkg/api"
	"github.com/ReolinkCameraAPI/reolinkapi/internal/pkg/enum"
	"github.com/ReolinkCameraAPI/reolinkapi/internal/pkg/models"
	"github.com/ReolinkCameraAPI/reolinkapi/pkg"
	"github.com/jarcoal/httpmock"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func registerMockGetRecordingEncoding() {
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

			encoding := &models.Encoding{
				Audio:   false,
				Channel: 0,
				MainStream: models.RecordingMainStream{
					BitRate:   enum.MAIN_BIT_RATE_1024.Value(),
					FrameRate: enum.MAIN_FRAME_RATE_20.Value(),
					Profile:   "Main",
					Size:      enum.MAIN_SIZE_2592_1944.Value(),
				},
				SubStream: models.RecordingSubStream{
					BitRate:   enum.SUB_BIT_RATE_192.Value(),
					FrameRate: enum.SUB_FRAME_RATE_10.Value(),
					Profile:   "High",
					Size:      enum.SUB_SIZE_640_360.Value(),
				},
			}

			generalData := map[string]interface{}{
				"cmd":  "GetEnc",
				"code": 0,
				"value": map[string]interface{}{
					"Enc": encoding,
				},
			}

			return httpmock.NewJsonResponse(200, []interface{}{generalData})
		},
	)
}

func registerMockGetRecordingAdvanced() {
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

			recording := &models.Recording{
				Channel:    0,
				Overwrite:  true,
				PostRecord: enum.POST_RECORD_SECONDS_30.Value(),
				PreRecord:  true,
				Schedule: models.Schedule{
					Enable: true,
					Table:  "000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
				},
			}

			generalData := map[string]interface{}{
				"cmd":  "GetRec",
				"code": 0,
				"value": map[string]interface{}{
					"Rec": recording,
				},
			}

			return httpmock.NewJsonResponse(200, []interface{}{generalData})
		},
	)
}

func registerMockSetRecordingEncoding() {
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

			var encoding *models.Encoding

			err = json.Unmarshal(reqData[0].Param["Enc"], &encoding)

			if err != nil {
				return httpmock.NewStringResponse(500, err.Error()), nil
			}

			log.Printf("encoding %v", encoding)

			generalData := map[string]interface{}{
				"cmd":  "SetEnc",
				"code": 0,
				"value": map[string]interface{}{
					"rspCode": 200,
				},
			}

			return httpmock.NewJsonResponse(200, []interface{}{generalData})
		},
	)
}

func TestRecordingMixin_GetRecordingEncoding(t *testing.T) {
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

	registerMockGetRecordingEncoding()

	encodingInfo, err := camera.API.GetRecordingEncoding()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("GetRecordingEncoding %v", encodingInfo)
}

func TestRecordingMixin_GetRecordingAdvanced(t *testing.T) {
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

	registerMockGetRecordingAdvanced()

	recordingInfo, err := camera.API.GetRecordingAdvanced()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("GetRecording %v", recordingInfo)
}

func TestRecordingMixin_SetRecordingEncoding(t *testing.T) {
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

	registerMockSetRecordingEncoding()

	recordingInfo, err := camera.API.SetRecordingEncoding(
		api.RecordingEncodingOptionMainBitRate(enum.MAIN_BIT_RATE_3072),
		api.RecordingEncodingOptionMainFrameRate(enum.MAIN_FRAME_RATE_8),
	)(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("SetRecordingEncoding %v", recordingInfo)
}
