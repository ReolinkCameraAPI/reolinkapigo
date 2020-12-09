package api

import (
	"encoding/json"
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/models"
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/network"
)

type SystemMixin struct{}

// Get the camera performance information
// See examples/responses/GetPerformance.json for example response data
func (sm *SystemMixin) GetPerformance() func(handler *network.RestHandler) (*models.DevicePerformanceInformation, error) {
	return func(handler *network.RestHandler) (*models.DevicePerformanceInformation, error) {
		payload := map[string]interface{}{
			"cmd":    "GetPerformance",
			"action": 0,
			"param":  map[string]interface{}{},
		}

		result, err := handler.Request("POST", payload, "GetPerformance", true)

		if err != nil {
			return nil, err
		}

		var devicePerformance *models.DevicePerformanceInformation

		err = json.Unmarshal(result.Value["GetPerformance"], &devicePerformance)

		if err != nil {
			return nil, err
		}

		return devicePerformance, nil
	}
}

// Get the camera device information
// See examples/responses/GetDevInfo.json for example response data
func (sm *SystemMixin) GetDeviceInformation() func(handler *network.RestHandler) (*models.DeviceInformation, error) {
	return func(handler *network.RestHandler) (*models.DeviceInformation, error) {
		payload := map[string]interface{}{
			"cmd":    "GetDevInfo",
			"action": 0,
			"param":  map[string]interface{}{},
		}

		result, err := handler.Request("POST", payload, "GetDevInfo", true)

		if err != nil {
			return nil, err
		}

		var deviceInfo *models.DeviceInformation

		err = json.Unmarshal(result.Value["DevInfo"], &deviceInfo)

		if err != nil {
			return nil, err
		}

		return deviceInfo, nil
	}
}

// Reboot the camera
func (sm *SystemMixin) RebootCamera() func(handler *network.RestHandler) (bool, error) {
	return func(handler *network.RestHandler) (bool, error) {
		payload := map[string]interface{}{
			"cmd":    "Reboot",
			"action": 0,
			"param":  map[string]interface{}{},
		}

		result, err := handler.Request("POST", payload, "Reboot", true)

		if err != nil {
			return false, err
		}

		var deviceRebootData *models.DeviceReboot
		// TODO: need to confirm this
		err = json.Unmarshal(result.Value["Reboot"], &deviceRebootData)

		if err != nil {
			return false, err
		}

		return true, nil
	}
}

// Get the camera DST information
// See examples/response/GetDSTInfo.json for example response data
func (sm *SystemMixin) GetDstInformation() func(handler *network.RestHandler) (*models.DstInformation,
	*models.TimeInformation, error) {
	return func(handler *network.RestHandler) (*models.DstInformation, *models.TimeInformation, error) {
		payload := map[string]interface{}{
			"cmd":    "GetTime",
			"action": 0,
			"param":  map[string]interface{}{},
		}

		result, err := handler.Request("POST", payload, "GetTime", true)

		if err != nil {
			return nil, nil, err
		}

		var dstData *models.DstInformation
		var timeData *models.TimeInformation

		err = json.Unmarshal(result.Value["Dst"], &dstData)

		if err != nil {
			return nil, nil, err
		}

		err = json.Unmarshal(result.Value["Time"], &timeData)

		if err != nil {
			return nil, nil, err
		}
		return dstData, timeData, nil
	}
}
