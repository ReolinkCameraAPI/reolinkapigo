package api

import "github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/enum"

type ImageMixin interface {
	SetAdvanceImageSettings(
		anti_flicker enum.AntiFlicker,
		exposure enum.Exposure,
		gain_min int,
		gain_max int,
		shutter_min int,
		shutter_max int,
		blue_gain int,
		red_gain int,
		white_balance string,
		day_night string,
	)
}
