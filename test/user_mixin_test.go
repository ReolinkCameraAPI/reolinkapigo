package test

import (
	"encoding/json"
	"github.com/ReolinkCameraAPI/reolinkapi/internal/pkg/enum"
	"github.com/ReolinkCameraAPI/reolinkapi/internal/pkg/models"
	"github.com/ReolinkCameraAPI/reolinkapi/pkg"
	"github.com/jarcoal/httpmock"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func registerMockGetOnlineUsers() {
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

			onlineUsers := []*models.User{
				{
					Level:             enum.USER_LEVEL_ADMIN.Value(),
					Username:          "admin",
					CanBeDisconnected: false,
					IP:                "192.168.1.100",
					SessionId:         1000,
				},
				{
					Level:             enum.USER_LEVEL_GUEST.Value(),
					Username:          "guest123",
					CanBeDisconnected: true,
					IP:                "192.168.1.101",
					SessionId:         1001,
				},
			}

			generalData := map[string]interface{}{
				"cmd":  "GetOnline",
				"code": 0,
				"value": map[string]interface{}{
					"User": onlineUsers,
				},
			}

			return httpmock.NewJsonResponse(200, []interface{}{generalData})

		},
	)
}

func registerMockGetUsers() {
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

			users := []*models.User{
				{
					Level:    enum.USER_LEVEL_ADMIN.Value(),
					Username: "admin",
				},
				{
					Level:    enum.USER_LEVEL_GUEST.Value(),
					Username: "guest123",
				},
			}

			generalData := map[string]interface{}{
				"cmd":  "GetUser",
				"code": 0,
				"value": map[string]interface{}{
					"User": users,
				},
			}

			return httpmock.NewJsonResponse(200, []interface{}{generalData})

		},
	)
}

func registerMockAddUser() {
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

			var user *models.User

			err = json.Unmarshal(reqData[0].Param["User"], &user)

			if err != nil {
				return httpmock.NewStringResponse(500, err.Error()), nil
			}

			log.Printf("adding user %v", user)

			generalData := map[string]interface{}{
				"cmd":  "AddUser",
				"code": 0,
				"value": map[string]interface{}{
					"rspCode": 200,
				},
			}

			return httpmock.NewJsonResponse(200, []interface{}{generalData})

		},
	)
}

func registerMockUpdateUserPassword() {
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

			var user *models.User

			err = json.Unmarshal(reqData[0].Param["User"], &user)

			if err != nil {
				return httpmock.NewStringResponse(500, err.Error()), nil
			}

			log.Printf("modifying user %v", user)

			generalData := map[string]interface{}{
				"cmd":  "ModifyUser",
				"code": 0,
				"value": map[string]interface{}{
					"rspCode": 200,
				},
			}

			return httpmock.NewJsonResponse(200, []interface{}{generalData})

		},
	)
}

func registerMockDeleteUser() {
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

			var user *models.User

			err = json.Unmarshal(reqData[0].Param["User"], &user)

			if err != nil {
				return httpmock.NewStringResponse(500, err.Error()), nil
			}

			log.Printf("modifying user %v", user)

			generalData := map[string]interface{}{
				"cmd":  "DelUser",
				"code": 0,
				"value": map[string]interface{}{
					"rspCode": 200,
				},
			}

			return httpmock.NewJsonResponse(200, []interface{}{generalData})

		},
	)
}

func TestUserMixin_GetOnlineUsers(t *testing.T) {
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

	registerMockGetOnlineUsers()

	onlineUsers, err := camera.API.GetOnlineUsers()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	data, _ := json.Marshal(onlineUsers)

	t.Logf("GetOnlineUsers %s", string(data))
}

func TestUserMixin_GetUsers(t *testing.T) {
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

	registerMockGetUsers()

	users, err := camera.API.GetUsers()(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	data, _ := json.Marshal(users)

	t.Logf("GetOnlineUsers %s", string(data))
}

func TestUserMixin_AddUser(t *testing.T) {
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

	registerMockAddUser()

	ok, err := camera.API.AddUser("user1234", "12345", enum.USER_LEVEL_GUEST)(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("AddUser %v", ok)
}

func TestUserMixin_UpdateUserPassword(t *testing.T) {
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

	registerMockUpdateUserPassword()

	ok, err := camera.API.UpdateUserPassword("user1234", "12345")(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("UpdateUserPassword %v", ok)
}

func TestUserMixin_DeleteUser(t *testing.T) {
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
	
	registerMockDeleteUser()

	ok, err := camera.API.DeleteUser("user1234")(camera.RestHandler)

	if err != nil {
		t.Error(err)
	}

	t.Logf("DeleteUser %v", ok)
}
