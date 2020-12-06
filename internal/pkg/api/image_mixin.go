package api

import (
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/enum"
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/network"
)

type ImageMixin struct {
}

func (im *ImageMixin) SetAdvanceImageSettings(
	antiFlicker enum.AntiFlicker,
	exposure enum.Exposure,
	gainMin int,
	gainMax int,
	shutterMin int,
	shutterMax int,
	blueGain int,
	redGain int,
	whiteBalance string,
	dayNight string,
) func(handler *network.RestHandler) (bool, error) {
	return func(handler *network.RestHandler) (bool, error) {
		// TODO: implement
		return false, nil
	}
}
