package api

import (
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/network"
)

type DisplayMixin struct {
}

func (dm *DisplayMixin) GetOSD() func(handler *network.RestHandler) (*network.OsdData, error) {
	return func(handler *network.RestHandler) (*network.OsdData, error) {
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

		osdData := result.Value.(*network.OsdData)
		return osdData, nil
	}
}

func (dm *DisplayMixin) GetMask() func(handler *network.RestHandler) (bool, error) {
	return func(handler *network.RestHandler) (bool, error) {
		// TODO: implement
		return false, nil
	}
}

func (dm *DisplayMixin) SetOSD() func(handler *network.RestHandler) (bool, error) {
	return func(handler *network.RestHandler) (bool, error) {
		// TODO: implement
		return false, nil
	}
}
