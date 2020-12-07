package api

import (
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/models"
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/network"
)

type DisplayMixin struct {
}

func (dm *DisplayMixin) GetOSD() func(handler *network.RestHandler) (*models.OsdData, error) {
	return func(handler *network.RestHandler) (*models.OsdData, error) {
		payload := map[string]interface{}{
			"cmd":    "GetOSD",
			"action": 1,
			"params": map[string]interface{}{
				"channel": 0,
			},
		}

		result, err := handler.Request("GET", payload, true)

		if err != nil {
			return nil, err
		}

		osdData := result.Value.(*models.OsdData)
		return osdData, nil
	}
}

func (dm *DisplayMixin) GetMask() func(handler *network.RestHandler) (*models.MaskData, error) {
	return func(handler *network.RestHandler) (*models.MaskData, error) {
		payload := map[string]interface{}{
			"cmd":    "GetMask",
			"action": 1,
			"param": map[string]interface{}{
				"channel": 0,
			},
		}

		result, err := handler.Request("GET", payload, true)

		if err != nil {
			return nil, err
		}

		maskData := result.Value.(*models.MaskData)

		return maskData, nil
	}
}

func (dm *DisplayMixin) SetOSD() func(handler *network.RestHandler) (bool, error) {
	return func(handler *network.RestHandler) (bool, error) {
		// TODO: implement
		return false, nil
	}
}
