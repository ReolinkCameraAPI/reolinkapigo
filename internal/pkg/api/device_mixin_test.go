package api

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

func TestDeviceMixin_GetHddInfo(t *testing.T) {

	httpmock.Activate()

	defer httpmock.DeactivateAndReset()

	loginData := &models.Login{Token: models.LoginToken{LeaseTime: 3600, Name: "12345"}}

	httpmock.RegisterResponder("POST", "https://127.0.0.1/cgi-bin/api.cgi",
		func(req *http.Request) (*http.Response, error) {

			type User struct {
				UserName string `json:"userName"`
				Password string `json:"password"`
			}

			type ReqData struct {
				Cmd    string `json:"cmd"`
				Action string `json:"action"`
				Param  map[string]json.RawMessage
			}

			// check the username and password
			var reqData ReqData

			data, err := ioutil.ReadAll(req.Body)

			if err != nil {
				return httpmock.NewStringResponse(500, err.Error()), nil
			}

			err = json.Unmarshal(data, &reqData)

			if err != nil {
				return httpmock.NewStringResponse(500, err.Error()), nil
			}

			var user User

			err = json.Unmarshal(reqData.Param["User"], &user)

			if err != nil {
				return httpmock.NewStringResponse(500, err.Error()), nil
			}

			var status int

			if user.UserName == "foo" && user.Password == "bar" {
				status = 200

			} else {
				status = 500
			}

			resp, err := httpmock.NewJsonResponse(status, loginData)

			if err != nil {
				return httpmock.NewStringResponse(500, err.Error()), nil
			}

			return resp, nil
		},
	)

	camera, err := pkg.NewCamera("foo", "bar", "127.0.0.1")

	if err != nil {
		t.Error(err)
	}

	if camera.RestHandler.Token == "12345" {
		t.Logf("login successful")
	}

	hddInfo, err := camera.API.GetHddInfo()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	log.Printf("%v", hddInfo)
}
