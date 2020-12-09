package api

import (
	"encoding/json"
	"fmt"
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/models"
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/network"
)

type DeviceMixin struct {
}

// Get the Camera's HDD information
// TODO: Better error messages
func (dm *DeviceMixin) GetHddInfo() func(handler *network.RestHandler) (*models.HddInfo, error) {
	return func(handler *network.RestHandler) (*models.HddInfo, error) {
		payload := map[string]interface{}{
			"cmd":    "GetHddInfo",
			"action": 0,
			"params": map[string]interface{}{},
		}

		result, err := handler.Request("POST", payload, "GetHddInfo", true)

		if err != nil {
			return nil, err
		}

		if result.Code == 0 {
			var hddInfoData *models.HddInfo
			err = json.Unmarshal(result.Value["HddInfo"], &hddInfoData)

			if err != nil {
				return nil, err
			}

			return hddInfoData, nil
		}

		return nil, fmt.Errorf("could not retrieve hdd info data")
	}
}

// Format the camera HDD.
// Default hddId: 0
// TODO: better error messages
func (dm *DeviceMixin) FormatHdd(hddId int) func(handler *network.RestHandler) (bool, error) {
	return func(handler *network.RestHandler) (bool, error) {
		payload := map[string]interface{}{
			"cmd":    "Format",
			"action": 0,
			"params": map[string]interface{}{
				"HddInfo": map[string]interface{}{
					"id": hddId,
				},
			},
		}

		result, err := handler.Request("POST", payload, "Format", true)

		if err != nil {
			return false, err
		}

		var respCode int

		err = json.Unmarshal(result.Value["rspCode"], &respCode)

		if err != nil {
			return false, err
		}

		if respCode == 200 {
			return true, nil
		}

		return false, fmt.Errorf("camera could not format hdd. camera responded with %v", result.Value)
	}
}
