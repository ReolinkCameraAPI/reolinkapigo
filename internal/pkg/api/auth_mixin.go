package api

import (
	"encoding/json"
	"fmt"
	"github.com/ReolinkCameraAPI/reolinkapigo/internal/pkg/models"
	"github.com/ReolinkCameraAPI/reolinkapigo/internal/pkg/network/rest"
	"log"
)

type AuthMixin struct {
	Username string
	Password string
	token    string
}

func (am *AuthMixin) Login() func(*rest.RestHandler) (bool, error) {
	return func(handler *rest.RestHandler) (bool, error) {

		payload := map[string]interface{}{
			"cmd":    "Login",
			"action": 0,
			"param": map[string]interface{}{
				"User": map[string]interface{}{
					"userName": am.Username,
					"password": am.Password,
				},
			},
		}

		result, err := handler.Request("POST", payload, "Login")

		if err != nil {
			return false, err
		}

		var tokenData *models.LoginToken

		err = json.Unmarshal(result.Value["Token"], &tokenData)

		if err != nil {
			return false, err
		}

		log.Printf("token data unmarshalled %v", tokenData)

		if tokenData == nil {
			return false, fmt.Errorf("login failed")
		}

		am.token = tokenData.Name
		handler.SetToken(tokenData.Name)

		return true, nil
	}
}

func (am *AuthMixin) Logout() func(handler *rest.RestHandler) (bool, error) {
	return func(handler *rest.RestHandler) (bool, error) {
		payload := map[string]interface{}{
			"cmd":    "Logout",
			"action": 0,
		}

		result, err := handler.Request("POST", payload, "Logout")

		if err != nil {
			return false, err
		}

		// Set the token
		if result.Code == 0 {
			var tokenData *models.LoginToken
			err = json.Unmarshal(result.Value["Token"], &tokenData)

			if err != nil {
				return false, err
			}

			if tokenData != nil {
				handler.SetToken(tokenData.Name)
			} else {
				return false, fmt.Errorf("token data could not be retrieved")
			}

			return true, nil
		}

		return false, fmt.Errorf("login failed")
	}
}

