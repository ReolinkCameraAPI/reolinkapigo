package api

import (
	"encoding/json"
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/enum"
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/models"
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/network"
)

type DisplayMixin struct {
}

// Get the camera's Osd information
func (dm *DisplayMixin) GetOSD() func(handler *network.RestHandler) (*models.Osd, error) {
	return func(handler *network.RestHandler) (*models.Osd, error) {
		payload := map[string]interface{}{
			"cmd":    "GetOsd",
			"action": 1,
			"param": map[string]interface{}{
				"channel": 0,
			},
		}

		result, err := handler.Request("POST", payload, "GetOsd", true)

		if err != nil {
			return nil, err
		}

		var osdData *models.Osd

		err = json.Unmarshal(result.Value["Osd"], &osdData)

		if err != nil {
			return nil, err
		}

		return osdData, nil
	}
}

// Get the camera's mask information
func (dm *DisplayMixin) GetMask() func(handler *network.RestHandler) (*models.MaskData, error) {
	return func(handler *network.RestHandler) (*models.MaskData, error) {
		payload := map[string]interface{}{
			"cmd":    "GetMask",
			"action": 1,
			"param": map[string]interface{}{
				"channel": 0,
			},
		}

		result, err := handler.Request("GET", payload, "GetMask", true)

		if err != nil {
			return nil, err
		}

		var maskData *models.MaskData

		err = json.Unmarshal(result.Value["Mask"], &maskData)

		if err != nil {
			return nil, err
		}

		return maskData, nil
	}
}

// Set the camera's Osd
func (dm *DisplayMixin) SetOSD(osdOption ...func(osd *models.Osd) error) func(handler *network.RestHandler) (bool,
	error) {

	osd := &models.Osd{
		BgColor: false,
		Channel: 0,
		OsdChannel: models.OsdChannel{
			Enable: true,
			Name:   "",
			Pos:    "Lower Right",
		},
		OsdTime: models.OsdTime{
			Enable: false,
			Pos:    "Lower Right",
		},
	}

	for _, op := range osdOption {
		err := op(osd)
		if err != nil {

		}
	}

	return func(handler *network.RestHandler) (bool, error) {
		payload := map[string]interface{}{
			"cmd":    "SetOsd",
			"action": 1,
			"param": map[string]interface{}{
				"Osd": map[string]interface{}{
					"bgcolor": osd.BgColor,
					"channel": osd.Channel,
					"osdChannel": map[string]interface{}{
						"enable": osd.OsdChannel.Enable,
						"name":   osd.OsdChannel.Name,
						"pos":    osd.OsdChannel.Pos,
					},
					"osdTime": map[string]interface{}{
						"enable": osd.OsdTime.Enable,
						"pos":    osd.OsdTime.Pos,
					},
				},
			},
		}

		result, err := handler.Request("GET", payload, "SetOsd", true)

		if err != nil {
			return false, err
		}

		var osdData *models.Osd

		err = json.Unmarshal(result.Value["Osd"], &osdData)

		if err != nil {
			return false, err
		}

		return true, nil
	}
}

// Set the OSD background color on or off
func SetOsdOptionBgColor(bgColor bool) func(dm *models.Osd) {
	return func(dm *models.Osd) {
		dm.BgColor = bgColor
	}
}

// Set the OSD channel
func SetOsdOptionsChannel(channel int) func(dm *models.Osd) {
	return func(dm *models.Osd) {
		dm.Channel = channel
	}
}

// Set the OSD channel on or off
func SetOsdOptionsChannelEnable(enable bool) func(dm *models.Osd) {
	return func(dm *models.Osd) {
		dm.OsdChannel.Enable = enable
	}
}

// Set the OSD channel name
func SetOsdOptionsChannelName(name string) func(dm *models.Osd) {
	return func(dm *models.Osd) {
		dm.OsdChannel.Name = name
	}
}

// Set the OSD channel position
func SetOsdOptionsChannelPos(position enum.OsdPosition) func(dm *models.Osd) {
	return func(dm *models.Osd) {
		dm.OsdChannel.Pos = position.Name()
	}
}

// Set the OSD time as on or off
func SetOsdOptionsTimeEnable(enable bool) func(dm *models.Osd) {
	return func(dm *models.Osd) {
		dm.OsdTime.Enable = enable
	}
}

// Set the OSD time position
func SetOsdOptionsTimePos(position enum.OsdPosition) func(dm *models.Osd) {
	return func(dm *models.Osd) {
		dm.OsdTime.Pos = position.Name()
	}
}

