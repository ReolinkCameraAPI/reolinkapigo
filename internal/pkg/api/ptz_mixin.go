package api

import (
	"encoding/json"
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/models"
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/network"
)

type PtzMixin struct {
}

type ptzOperationOptions struct {
	Operation string
	Speed     *int
	Index     *int
}

type ptzPresetOptions struct {
	Index int
	Name  string
}

type OptionPtzOperation func(*ptzOperationOptions)

type OptionPtzPreset func(*ptzPresetOptions)

// helper function for ptz presets
func ptzPreset(enable bool, preset int, name string) interface{} {
	return map[string]interface{}{
		"cmd":    "SetPtzPreset",
		"action": 0,
		"param": map[string]interface{}{
			"channel": 0,
			"enable":  enable,
			"id":      preset,
			"name":    name,
		},
	}
}

// helper function for ptz operations
func ptzOperation(ptzOperation *ptzOperationOptions) interface{} {

	param := map[string]interface{}{
		"channel": 0,
		"op":      ptzOperation.Operation,
	}

	if ptzOperation.Index != nil {
		param["index"] = ptzOperation.Index
	}

	if ptzOperation.Speed != nil {
		param["speed"] = ptzOperation.Speed
	}

	return map[string]interface{}{
		"cmd":    "PtzCtrl",
		"action": 0,
		"param":  param,
	}
}

// Moves the camera to the specified preset
// The preset index and speed is optional and will fallback to defaults
// Defaults:
// index: 1
// speed: 60
func (pm *PtzMixin) GoToPreset(ptzOptions ...OptionPtzOperation) func(handler *network.RestHandler) (
	*models.PtzOperation, error) {
	speed := 60
	index := 1

	ptzPreset := &ptzOperationOptions{
		Operation: "ToPos",
		Speed:     &speed,
		Index:     &index,
	}

	for _, op := range ptzOptions {
		op(ptzPreset)
	}

	return func(handler *network.RestHandler) (*models.PtzOperation, error) {
		payload := ptzOperation(ptzPreset)
		result, err := handler.Request("POST", payload, "PtzCtrl", true)

		if err != nil {
			return nil, err
		}

		var ptzOperation *models.PtzOperation
		// TODO: will need to confirm this
		err = json.Unmarshal(result.Value["Ptz"], ptzOperation)

		if err != nil {
			return nil, err
		}

		return ptzOperation, nil
	}
}

// Create a new preset at the current camera position
// The preset index and name is optional and will fallback to defaults
// Defaults:
// index: 1
// name: pos1
func (pm *PtzMixin) AddPreset(ptzOptions ...OptionPtzPreset) func(handler *network.RestHandler) (bool, error) {
	presetOptions := &ptzPresetOptions{
		Index: 1,
		Name:  "pos1",
	}

	for _, op := range ptzOptions {
		op(presetOptions)
	}

	return func(handler *network.RestHandler) (bool, error) {
		payload := ptzPreset(true, presetOptions.Index, presetOptions.Name)

		result, err := handler.Request("POST", payload, "PtzCtrl", true)

		if err != nil {
			return false, err
		}

		var ptzPreset *models.PtzPreset

		// TODO: will need to confirm this
		err = json.Unmarshal(result.Value["Ptz"], &ptzPreset)

		if err != nil {
			return false, err
		}

		return true, nil
	}
}

// Remove the specified preset
// The preset index and name is optional and will fallback to defaults
// Defaults:
// index: 1
// name: pos1
func (pm *PtzMixin) RemovePreset(ptzOptions ...OptionPtzPreset) func(handler *network.RestHandler) (bool, error) {

	presetOptions := &ptzPresetOptions{
		Index: 1,
		Name:  "pos1",
	}

	for _, op := range ptzOptions {
		op(presetOptions)
	}

	return func(handler *network.RestHandler) (bool, error) {
		payload := ptzPreset(false, presetOptions.Index, presetOptions.Name)

		result, err := handler.Request("POST", payload, "PtzPreset", true)

		if err != nil {
			return false, err
		}

		var ptzPreset *models.PtzPreset

		err = json.Unmarshal(result.Value["Ptz"], ptzPreset)

		if err != nil {
			return false, err
		}

		return true, nil
	}
}

// Move the camera to the right
// The operation speed is optional and will fallback to defaults. Other operations will be ignored.
// Defaults:
// speed: 25
func (pm *PtzMixin) MoveRight(ptzOptions ...OptionPtzOperation) func(handler *network.RestHandler) (bool,
	error) {

	speed := 25

	ptzOperations := &ptzOperationOptions{
		Operation: "Right",
		Speed:     &speed,
		Index:     nil,
	}

	for _, op := range ptzOptions {
		op(ptzOperations)
	}

	return func(handler *network.RestHandler) (bool, error) {
		// set the index to nil in case the user passes an option for it
		ptzOperations.Index = nil
		payload := ptzOperation(ptzOperations)

		result, err := handler.Request("POST", payload, "PtzCtrl", true)

		if err != nil {
			return false, err
		}

		var ptzOperation *models.PtzOperation

		err = json.Unmarshal(result.Value["Ptz"], ptzOperation)

		if err != nil {
			return false, err
		}

		return false, nil
	}
}

// Move the camera to the right Up
// The operation speed is optional and will fallback to defaults. Other operations will be ignored.
// Defaults:
// speed: 25
func (pm *PtzMixin) MoveRightUp(ptzOptions ...OptionPtzOperation) func(handler *network.RestHandler) (bool,
	error) {

	speed := 25

	ptzOperations := &ptzOperationOptions{
		Operation: "RightUp",
		Speed:     &speed,
		Index:     nil,
	}

	for _, op := range ptzOptions {
		op(ptzOperations)
	}

	return func(handler *network.RestHandler) (bool, error) {
		// set the index to nil in case the user passes an option for it
		ptzOperations.Index = nil
		payload := ptzOperation(ptzOperations)

		result, err := handler.Request("POST", payload, "PtzCtrl", true)

		if err != nil {
			return false, err
		}

		var ptzOperation *models.PtzOperation

		err = json.Unmarshal(result.Value["Ptz"], ptzOperation)

		if err != nil {
			return false, err
		}

		return false, nil
	}
}

// Move the camera to the right Down
// The operation speed is optional and will fallback to defaults. Other operations will be ignored.
// Defaults:
// speed: 25
func (pm *PtzMixin) MoveRightDown(ptzOptions ...OptionPtzOperation) func(handler *network.RestHandler) (bool,
	error) {

	speed := 25

	ptzOperations := &ptzOperationOptions{
		Operation: "RightDown",
		Speed:     &speed,
		Index:     nil,
	}

	for _, op := range ptzOptions {
		op(ptzOperations)
	}

	return func(handler *network.RestHandler) (bool, error) {
		// set the index to nil in case the user passes an option for it
		ptzOperations.Index = nil
		payload := ptzOperation(ptzOperations)

		result, err := handler.Request("POST", payload, "PtzCtrl", true)

		if err != nil {
			return false, err
		}

		var ptzOperation *models.PtzOperation

		err = json.Unmarshal(result.Value["Ptz"], ptzOperation)

		if err != nil {
			return false, err
		}

		return false, nil
	}
}

// Move the camera to the right Left
// The operation speed is optional and will fallback to defaults. Other operations will be ignored.
// Defaults:
// speed: 25
func (pm *PtzMixin) MoveLeft(ptzOptions ...OptionPtzOperation) func(handler *network.RestHandler) (bool,
	error) {

	speed := 25

	ptzOperations := &ptzOperationOptions{
		Operation: "Left",
		Speed:     &speed,
		Index:     nil,
	}

	for _, op := range ptzOptions {
		op(ptzOperations)
	}

	return func(handler *network.RestHandler) (bool, error) {
		// set the index to nil in case the user passes an option for it
		ptzOperations.Index = nil
		payload := ptzOperation(ptzOperations)

		result, err := handler.Request("POST", payload, "PtzCtrl", true)

		if err != nil {
			return false, err
		}

		var ptzOperation *models.PtzOperation

		err = json.Unmarshal(result.Value["Ptz"], ptzOperation)

		if err != nil {
			return false, err
		}

		return false, nil
	}
}

// Move the camera to the left Up
// The operation speed is optional and will fallback to defaults. Other operations will be ignored.
// Defaults:
// speed: 25
func (pm *PtzMixin) MoveLeftUp(ptzOptions ...OptionPtzOperation) func(handler *network.RestHandler) (bool,
	error) {

	speed := 25

	ptzOperations := &ptzOperationOptions{
		Operation: "LeftUp",
		Speed:     &speed,
		Index:     nil,
	}

	for _, op := range ptzOptions {
		op(ptzOperations)
	}

	return func(handler *network.RestHandler) (bool, error) {
		// set the index to nil in case the user passes an option for it
		ptzOperations.Index = nil
		payload := ptzOperation(ptzOperations)

		result, err := handler.Request("POST", payload, "PtzCtrl", true)

		if err != nil {
			return false, err
		}

		var ptzOperation *models.PtzOperation

		err = json.Unmarshal(result.Value["Ptz"], ptzOperation)

		if err != nil {
			return false, err
		}

		return false, nil
	}
}

// Move the camera to the left Down
// The operation speed is optional and will fallback to defaults. Other operations will be ignored.
// Defaults:
// speed: 25
func (pm *PtzMixin) MoveLeftDown(ptzOptions ...OptionPtzOperation) func(handler *network.RestHandler) (bool,
	error) {

	speed := 25

	ptzOperations := &ptzOperationOptions{
		Operation: "LeftDown",
		Speed:     &speed,
		Index:     nil,
	}

	for _, op := range ptzOptions {
		op(ptzOperations)
	}

	return func(handler *network.RestHandler) (bool, error) {
		// set the index to nil in case the user passes an option for it
		ptzOperations.Index = nil
		payload := ptzOperation(ptzOperations)

		result, err := handler.Request("POST", payload, "PtzCtrl", true)

		if err != nil {
			return false, err
		}

		var ptzOperation *models.PtzOperation

		err = json.Unmarshal(result.Value["Ptz"], ptzOperation)

		if err != nil {
			return false, err
		}

		return false, nil
	}
}

// Move the camera to the up
// The operation speed is optional and will fallback to defaults. Other operations will be ignored.
// Defaults:
// speed: 25
func (pm *PtzMixin) MoveUp(ptzOptions ...OptionPtzOperation) func(handler *network.RestHandler) (bool,
	error) {

	speed := 25

	ptzOperations := &ptzOperationOptions{
		Operation: "Up",
		Speed:     &speed,
		Index:     nil,
	}

	for _, op := range ptzOptions {
		op(ptzOperations)
	}

	return func(handler *network.RestHandler) (bool, error) {
		// set the index to nil in case the user passes an option for it
		ptzOperations.Index = nil
		payload := ptzOperation(ptzOperations)

		result, err := handler.Request("POST", payload, "PtzCtrl", true)

		if err != nil {
			return false, err
		}

		var ptzOperation *models.PtzOperation

		err = json.Unmarshal(result.Value["Ptz"], ptzOperation)

		if err != nil {
			return false, err
		}

		return false, nil
	}
}

// Move the camera to the down
// The operation speed is optional and will fallback to defaults. Other operations will be ignored.
// Defaults:
// speed: 25
func (pm *PtzMixin) MoveDown(ptzOptions ...OptionPtzOperation) func(handler *network.RestHandler) (bool,
	error) {

	speed := 25

	ptzOperations := &ptzOperationOptions{
		Operation: "Down",
		Speed:     &speed,
		Index:     nil,
	}

	for _, op := range ptzOptions {
		op(ptzOperations)
	}

	return func(handler *network.RestHandler) (bool, error) {
		// set the index to nil in case the user passes an option for it
		ptzOperations.Index = nil
		payload := ptzOperation(ptzOperations)

		result, err := handler.Request("POST", payload, "PtzCtrl", true)

		if err != nil {
			return false, err
		}

		var ptzOperation *models.PtzOperation

		err = json.Unmarshal(result.Value["Ptz"], ptzOperation)

		if err != nil {
			return false, err
		}

		return false, nil
	}
}

// Stops the cameras current action
func (pm *PtzMixin) StopPtz() func(handler *network.RestHandler) (bool,
	error) {
	return func(handler *network.RestHandler) (bool, error) {
		ptzOperations := &ptzOperationOptions{
			Operation: "Stop",
			Speed:     nil,
			Index:     nil,
		}

		payload := ptzOperation(ptzOperations)

		result, err := handler.Request("POST", payload, "PtzCtrl", true)

		if err != nil {
			return false, err
		}

		var ptzOperation *models.PtzOperation

		err = json.Unmarshal(result.Value["Ptz"], ptzOperation)

		if err != nil {
			return false, err
		}

		return false, nil
	}
}

// Move the camera in a clockwise rotation
func (pm *PtzMixin) AutoMovement() func(handler *network.RestHandler) (bool, error) {
	return func(handler *network.RestHandler) (bool, error) {
		ptzOperations := &ptzOperationOptions{
			Operation: "Auto",
			Speed:     nil,
			Index:     nil,
		}

		payload := ptzOperation(ptzOperations)

		result, err := handler.Request("POST", payload, "PtzCtrl", true)

		if err != nil {
			return false, err
		}

		var ptzOperation *models.PtzOperation

		err = json.Unmarshal(result.Value["Ptz"], ptzOperation)

		if err != nil {
			return false, err
		}

		return false, nil
	}
}

// Set the Ptz Operation Speed
func PtzOptionOpsSpeed(speed int) OptionPtzOperation {
	return func(p *ptzOperationOptions) {
		p.Speed = &speed
	}
}

// Set the Ptz Operation Index
func PtzOptionOpsIndex(index int) OptionPtzOperation {
	return func(p *ptzOperationOptions) {
		p.Index = &index
	}
}

// Set the Ptz Preset Index
func PtzOptionPresetIndex(index int) OptionPtzPreset {
	return func(p *ptzPresetOptions) {
		p.Index = index
	}
}

// Set the Ptz Preset Name
func PtzOptionsPresetName(name string) OptionPtzPreset {
	return func(p *ptzPresetOptions) {
		p.Name = name
	}
}
