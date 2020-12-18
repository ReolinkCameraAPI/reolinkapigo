package api

import (
	"encoding/json"
	"fmt"
	"github.com/ReolinkCameraAPI/reolinkapigo/internal/pkg/network/rest"
)

type ZoomFocusMixin struct{}

type zoom struct {
	operation string
	speed     *int
}

type focus struct {
	operation string
	speed     *int
}

type OptionZoomOperation func(*zoom)

type OptionFocusOperation func(*focus)

// zoom helper
func zoomOperation(zoomOperation *zoom) interface{} {

	param := map[string]interface{}{
		"channel": 0,
		"op":      zoomOperation.operation,
	}

	if zoomOperation.speed != nil {
		param["speed"] = zoomOperation.speed
	}

	return map[string]interface{}{
		"cmd":    "PtzCtrl",
		"action": 0,
		"param":  param,
	}
}

// focus helper
func focusOperation(focusOperation *focus) interface{} {
	param := map[string]interface{}{
		"channel": 0,
		"op":      focusOperation.operation,
	}

	if focusOperation.speed != nil {
		param["speed"] = focusOperation.speed
	}

	return map[string]interface{}{
		"cmd":    "PtzCtrl",
		"action": 0,
		"param":  param,
	}
}

// Zoom in with the camera with optional parameters.
// Defaults:
// speed: 60
func (zfm *ZoomFocusMixin) StartZoomingIn(zoomOptions ...OptionZoomOperation) func(handler *rest.RestHandler) (bool,
	error) {

	speed := 60

	zoomOps := &zoom{
		operation: "ZoomInc",
		speed:     &speed,
	}

	for _, op := range zoomOptions {
		op(zoomOps)
	}

	return func(handler *rest.RestHandler) (bool, error) {
		payload := zoomOperation(zoomOps)

		result, err := handler.Request("POST", payload, "PtzCtrl")

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

		return false, fmt.Errorf("camera could not zoom in. camera responded with %v", result.Value)
	}
}

// Zoom out with the camera with optional parameters.
// Default:
// speed: 60
func (zfm *ZoomFocusMixin) StartZoomingOut(zoomOptions ...OptionZoomOperation) func(handler *rest.RestHandler) (bool,
	error) {

	speed := 60

	zoomOps := &zoom{
		operation: "ZoomDec",
		speed:     &speed,
	}

	for _, op := range zoomOptions {
		op(zoomOps)
	}

	return func(handler *rest.RestHandler) (bool, error) {
		payload := zoomOperation(zoomOps)

		result, err := handler.Request("POST", payload, "PtzCtrl")

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

		return false, fmt.Errorf("camera could not zoom out. camera responded with %v", result.Value)
	}
}

// Stop zooming
func (zfm *ZoomFocusMixin) StopZooming() func(handler *rest.RestHandler) (bool, error) {
	zoomOps := &zoom{
		operation: "Stop",
		speed:     nil,
	}
	return func(handler *rest.RestHandler) (bool, error) {
		payload := zoomOperation(zoomOps)

		result, err := handler.Request("POST", payload, "PtzCtrl")

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

		return false, fmt.Errorf("camera could not stop zoom operation. camera responded with %v", result.Value)
	}
}

// Focus in with the camera with optional parameters.
// Defaults:
// speed: 32
func (zfm *ZoomFocusMixin) StartFocusingIn(focusOptions ...OptionFocusOperation) func(handler *rest.RestHandler) (
	bool, error) {
	speed := 32

	focusOps := &focus{
		operation: "FocusInc",
		speed:     &speed,
	}

	for _, op := range focusOptions {
		op(focusOps)
	}

	return func(handler *rest.RestHandler) (bool, error) {
		payload := focusOperation(focusOps)

		result, err := handler.Request("POST", payload, "PtzCtrl")

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

		return false, fmt.Errorf("camera could not focus in. camera responded with %v", result.Value)
	}
}

// Focus out with the camera with optional parameters.
// Defaults:
// speed: 32
func (zfm *ZoomFocusMixin) StartFocusingOut(focusOptions ...OptionFocusOperation) func(handler *rest.RestHandler) (
	bool, error) {

	speed := 32

	focusOps := &focus{
		operation: "FocusDec",
		speed:     &speed,
	}

	for _, op := range focusOptions {
		op(focusOps)
	}

	return func(handler *rest.RestHandler) (bool, error) {
		payload := focusOperation(focusOps)

		result, err := handler.Request("POST", payload, "PtzCtrl")

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

		return false, fmt.Errorf("camera could not focus out. camera responded with %v", result.Value)
	}
}

func (zfm *ZoomFocusMixin) StopFocusing() func(handler *rest.RestHandler) (bool, error) {

	focusOps := &focus{
		operation: "Stop",
		speed:     nil,
	}

	return func(handler *rest.RestHandler) (bool, error) {

		payload := focusOperation(focusOps)

		result, err := handler.Request("POST", payload, "PtzCtrl")

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

		return false, fmt.Errorf("camera could not stop focus operation. camera responded with %v", result.Value)

	}
}

// Set the zoom speed
// default: 60
func ZoomOptionSpeed(speed int) OptionZoomOperation {
	return func(z *zoom) {
		z.speed = &speed
	}
}

// Set the focus speed
// default: 32
func FocusOptionSpeed(speed int) OptionFocusOperation {
	return func(f *focus) {
		f.speed = &speed
	}
}
