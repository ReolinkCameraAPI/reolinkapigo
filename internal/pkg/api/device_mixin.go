package api

import (
	"fmt"
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/models"
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/network"
)

type DeviceMixin struct {
}


// Get the Camera's HDD information
// TODO: Better error messages
func (dm *DeviceMixin) GetHddInfo() func(handler *network.RestHandler) (*models.GetHddInfoData, error) {
	return func(handler *network.RestHandler) (*models.GetHddInfoData, error) {
		payload := map[string]interface{}{
			"cmd":    "GetHddInfo",
			"action": 0,
			"params": map[string]interface{}{},
		}

		result, err := handler.Request("GET", payload, true)

		if err != nil {
			return nil, err
		}

		if result.Code == 0 {
			hddInfoData := result.Value.(*models.GetHddInfoData)
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

		result, err := handler.Request("GET", payload, true)

		if err != nil {
			return false, err
		}

		formatHdd := result.Value.(*models.FormatHddData)

		if formatHdd.RspCode == 200 {
			return true, nil
		}

		return false, fmt.Errorf("could not format camera hdd")
	}
}
