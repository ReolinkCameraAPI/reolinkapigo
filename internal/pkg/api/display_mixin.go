package api

import "github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/network"

type DisplayMixin struct {
}

func (dm *DisplayMixin) GetOSD() func(handler *network.RestHandler) (bool, error) {
	return func(handler *network.RestHandler) (bool, error) {
		// TODO: implement
		return false, nil
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
