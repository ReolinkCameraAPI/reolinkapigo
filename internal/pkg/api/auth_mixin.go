package api

import (
	"fmt"
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/models"
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/network"
)

type AuthMixin struct {
}

func (am *AuthMixin) Login() func(handler *network.RestHandler) (bool, error) {
	return func(handler *network.RestHandler) (bool, error) {

		payload := map[string]interface{}{
			"cmd":    "Login",
			"action": 0,
			"params": map[string]interface{}{
				"User": map[string]interface{}{
					"userName": handler.Username,
					"password": handler.Password,
				},
			},
		}

		result, err := handler.Request("POST", payload, false)

		if err != nil {
			return false, err
		}

		// Set the token
		if result.Code == 0 {
			tokenData := result.Value.(*models.LoginData)
			handler.SetToken(tokenData.Token.Name)

			return true, nil
		}

		return false, fmt.Errorf("login failed")
	}
}
