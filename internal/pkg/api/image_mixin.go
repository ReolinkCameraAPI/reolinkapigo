package api

import (
	"encoding/json"
	"fmt"
	"github.com/ReolinkCameraAPI/reolinkapigo/internal/pkg/enum"
	"github.com/ReolinkCameraAPI/reolinkapigo/internal/pkg/network"
)

type ImageMixin struct {
}

type imageAdvanced struct {
	antiFlicker  string
	exposure     string
	gainMin      int
	gainMax      int
	shutterMin   int
	shutterMax   int
	blueGain     int
	redGain      int
	whiteBalance string
	dayNight     string
	backLight    string
	blc          int
	drc          int
	rotation     int
	mirroring    int
	nr3d         int
}

type image struct {
	brightness int
	contrast   int
	hue        int
	saturation int
	sharpness  int
}

type OptionAdvancedImageSetting func(*imageAdvanced)

type OptionImageSetting func(*image)

// Set the Advanced Image setting. Parameters are optional and will fallback to defaults.
// Defaults:
// antiFlicker:  "Outdoor",
// exposure:     "Auto",
// gainMin:      1,
// gainMax:      62,
// shutterMin:   1,
// shutterMax:   125,
// blueGain:     128,
// redGain:      128,
// whiteBalance: "Auto",
// dayNight:     "Auto",
// backLight:    "DynamicRangeControl",
// blc:          128,
// drc:          128,
// rotation:     0,
// mirroring:    0,
// nr3d:         1,
func (im *ImageMixin) SetAdvanceImageSettings(imageAdvancedOptions ...OptionAdvancedImageSetting) func(handler *network.RestHandler) (bool,
	error) {

	ias := &imageAdvanced{
		antiFlicker:  "Outdoor",
		exposure:     "Auto",
		gainMin:      1,
		gainMax:      62,
		shutterMin:   1,
		shutterMax:   125,
		blueGain:     128,
		redGain:      128,
		whiteBalance: "Auto",
		dayNight:     "Auto",
		backLight:    "DynamicRangeControl",
		blc:          128,
		drc:          128,
		rotation:     0,
		mirroring:    0,
		nr3d:         1,
	}

	for _, op := range imageAdvancedOptions {
		op(ias)
	}

	return func(handler *network.RestHandler) (bool, error) {

		payload := map[string]interface{}{
			"cmd":    "SetIsp",
			"action": 0,
			"param": map[string]interface{}{
				"Isp": map[string]interface{}{
					"channel":     0,
					"antiFlicker": ias.antiFlicker,
					"exposure":    ias.exposure,
					"gain": map[string]interface{}{
						"min": ias.gainMin,
						"max": ias.gainMax,
					},
					"shutter": map[string]interface{}{
						"min": ias.shutterMin,
						"max": ias.shutterMax,
					},
					"blueGain":     ias.blueGain,
					"redGain":      ias.redGain,
					"whiteBalance": ias.whiteBalance,
					"dayNight":     ias.dayNight,
					"backLight":    ias.backLight,
					"blc":          ias.blc,
					"drc":          ias.drc,
					"rotation":     ias.rotation,
					"mirroring":    ias.mirroring,
					"nr3d":         ias.nr3d,
				},
			},
		}

		result, err := handler.Request("POST", payload, "SetIsp", true)

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

		return false, fmt.Errorf("camera could not set advanced image settings. camera responded with %v", result.Value)
	}
}

// Set the Image Settings. Parameters are optional and will fallback to defautls.
// Defaults:
// brightness: 128,
// contrast:   62,
// hue:        1,
// saturation: 125,
// sharpness:  128,
func (im *ImageMixin) SetImageSettings(imageOptions ...OptionImageSetting) func(handler *network.RestHandler) (bool,
	error) {

	img := &image{
		brightness: 128,
		contrast:   62,
		hue:        1,
		saturation: 125,
		sharpness:  128,
	}

	for _, op := range imageOptions {
		op(img)
	}

	return func(handler *network.RestHandler) (bool, error) {
		payload := map[string]interface{}{
			"cmd":    "SetImage",
			"action": 0,
			"param": map[string]interface{}{
				"Image": map[string]interface{}{
					"bright":     img.brightness,
					"channel":    0,
					"contrast":   img.contrast,
					"hue":        img.hue,
					"saturation": img.saturation,
					"sharpen":    img.sharpness,
				},
			},
		}

		result, err := handler.Request("POST", payload, "SetImage", true)

		if err != nil {
			return false, err
		}

		var respCode int

		err = json.Unmarshal(result.Value["rspCode"], &respCode)

		if err != nil {
			return false, nil
		}

		if respCode == 200 {
			return true, nil
		}

		return false, fmt.Errorf("camera could not set image settings. camera responded with %v", result.Value)
	}
}

// Set Image Brightness
// Default: 128
func ImageOptionBrightness(brightness int) OptionImageSetting {
	return func(i *image) {
		i.brightness = brightness
	}
}

// Set Image Contrast
// Default: 62
func ImageOptionContrast(contrast int) OptionImageSetting {
	return func(i *image) {
		i.contrast = contrast
	}
}

// Set Image Hue
// Default: 1
func ImageOptionHue(hue int) OptionImageSetting {
	return func(i *image) {
		i.hue = hue
	}
}

// Set Image Saturation
// Default: 125
func ImageOptionSaturation(saturation int) OptionImageSetting {
	return func(i *image) {
		i.saturation = saturation
	}
}

// Set Image Sharpness
// Default: 128
func ImageOptionSharpness(sharpness int) OptionImageSetting {
	return func(i *image) {
		i.sharpness = sharpness
	}
}

// Set the anti flicker value
// Default: Outdoor
func ImageAdvancedOptionAntiFlicker(antiFlicker enum.AntiFlicker) OptionAdvancedImageSetting {
	return func(advanced *imageAdvanced) {
		advanced.antiFlicker = antiFlicker.Value()
	}
}

// Set the exposure value
// Default: Auto
func ImageAdvancedOptionExposure(exposure enum.Exposure) OptionAdvancedImageSetting {
	return func(advanced *imageAdvanced) {
		advanced.exposure = exposure.Value()
	}
}

// Set the gain min value
// Default: 1
func ImageAdvancedOptionGainMin(gainMin int) OptionAdvancedImageSetting {
	return func(advanced *imageAdvanced) {
		advanced.gainMin = gainMin
	}
}

// Set the gain max value
// Default: 62
func ImageAdvancedOptionGainMax(gainMax int) OptionAdvancedImageSetting {
	return func(advanced *imageAdvanced) {
		advanced.gainMax = gainMax
	}
}

// Set the shutter min value
// Default: 1
func ImageAdvancedOptionShutterMin(shutterMin int) OptionAdvancedImageSetting {
	return func(advanced *imageAdvanced) {
		advanced.shutterMin = shutterMin
	}
}

// Set the shutter max value
// Default: 125
func ImageAdvancedOptionShutterMax(shutterMax int) OptionAdvancedImageSetting {
	return func(advanced *imageAdvanced) {
		advanced.shutterMax = shutterMax
	}
}

// Set the blue gain value
// Default: 128
func ImageAdvancedOptionBlueGain(blueGain int) OptionAdvancedImageSetting {
	return func(advanced *imageAdvanced) {
		advanced.blueGain = blueGain
	}
}

// Set the red gain value
// Default: 128
func ImageAdvancedOptionRedGain(redGain int) OptionAdvancedImageSetting {
	return func(advanced *imageAdvanced) {
		advanced.redGain = redGain
	}
}

// Set the white balance value
// Default: Auto
func ImageAdvancedOptionWhiteBalance(whiteBalance enum.WhiteBalance) OptionAdvancedImageSetting {
	return func(advanced *imageAdvanced) {
		advanced.whiteBalance = whiteBalance.Value()
	}
}

// Set the day night value
// Default: Auto
func ImageAdvancedOptionDayNight(dayNight enum.DayNight) OptionAdvancedImageSetting {
	return func(advanced *imageAdvanced) {
		advanced.dayNight = dayNight.Value()
	}
}

// Set the backlight value
// Default: DynamicRangeControl
func ImageAdvancedOptionBacklight(backlight enum.Backlight) OptionAdvancedImageSetting {
	return func(advanced *imageAdvanced) {
		advanced.backLight = backlight.Value()
	}
}

// Set the blc value
// Default: 128
func ImageAdvancedOptionBlc(blc int) OptionAdvancedImageSetting {
	return func(advanced *imageAdvanced) {
		advanced.blc = blc
	}
}

// Set the drc value
// Default: 128
func ImageAdvancedOptionDrc(drc int) OptionAdvancedImageSetting {
	return func(advanced *imageAdvanced) {
		advanced.drc = drc
	}
}

// Set the rotation value
// Default: 0
func ImageAdvancedOptionRotation(rotation enum.Rotation) OptionAdvancedImageSetting {
	return func(advanced *imageAdvanced) {
		advanced.rotation = rotation.Value()
	}
}

// Set the mirroring value
// Default: 0
func ImageAdvancedOptionMirroring(mirroring int) OptionAdvancedImageSetting {
	return func(advanced *imageAdvanced) {
		advanced.mirroring = mirroring
	}
}

// Set the nr3d value
// Default: 0
func ImageAdvancedOptionNr3d(nr3d int) OptionAdvancedImageSetting {
	return func(advanced *imageAdvanced) {
		advanced.nr3d = nr3d
	}
}
