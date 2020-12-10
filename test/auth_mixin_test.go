package test

import (
	"encoding/json"
	"github.com/ReolinkCameraAPI/reolink-go-api/pkg"
	"github.com/jarcoal/httpmock"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestAuthMixin_Login(t *testing.T) {
	httpmock.Activate()

	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "http://127.0.0.1/cgi-bin/api.cgi",
		func(req *http.Request) (*http.Response, error) {

			type User struct {
				UserName string `json:"userName"`
				Password string `json:"password"`
			}

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

			var user User

			err = json.Unmarshal(reqData[0].Param["User"], &user)

			if err != nil {
				return httpmock.NewStringResponse(500, err.Error()), nil
			}

			var status int

			if user.UserName != "foo" || user.Password != "bar" {
				return httpmock.NewStringResponse(500, "username or password incorrect"), nil
			}

			loginData := map[string]interface{}{
				"Token": map[string]interface{}{
					"Name":     "12345",
					"LeaseTime": 3600,
				},
			}

			generalData := map[string]interface{}{
				"cmd":   "Login",
				"code":  0,
				"value": loginData,
			}

			return httpmock.NewJsonResponse(status, []interface{}{generalData})
		},
	)

	camera, err := pkg.NewCamera("foo", "bar", "127.0.0.1")

	if err != nil {
		t.Error(err)
	}

	if camera.RestHandler.Token == "12345" {
		t.Logf("login successful")
	}

}
