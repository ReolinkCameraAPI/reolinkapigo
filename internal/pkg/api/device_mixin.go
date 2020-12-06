package api

import "github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/network"

type DeviceMixin struct {
}

func (dm *DeviceMixin) GetHddInfo() func(handler *network.RestHandler) (bool, error) {
	return func(handler *network.RestHandler) (bool, error) {
		// TODO: implement
		return false, nil
	}
}

func (dm *DeviceMixin) FormatHdd() func(handler *network.RestHandler) (bool, error) {
	return func(handler *network.RestHandler) (bool, error) {
		// TODO: implement
		return false, nil
	}
}
