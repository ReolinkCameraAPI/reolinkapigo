package api

import (
	"encoding/json"
	"fmt"
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/network"
)

type AuthMixin struct {
	Username string
	Password string
	Token    string
}

func (am *AuthMixin) Login() func(handler *network.RestHandler) (bool, error) {
	payload := map[string]interface{}{
		"cmd":    "Login",
		"action": 0,
		"params": map[string]interface{}{
			"User": map[string]interface{}{
				"userName": am.Username,
				"password": am.Password,
			},
		},
	}

	data, err := json.Marshal([]interface{}{payload})

	return func(handler *network.RestHandler) (bool, error) {

		if err != nil {
			return false, err
		}

		resp, err := handler.Request("POST", data)

		if err != nil {
			return false, err
		}

		if resp.StatusCode == 200 {
			return true, nil
		}

		return false, fmt.Errorf("login failed")
	}
}
