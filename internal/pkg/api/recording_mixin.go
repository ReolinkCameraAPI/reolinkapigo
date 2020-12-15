package api

import (
	"encoding/json"
	"fmt"
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/enum"
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/models"
	"github.com/ReolinkCameraAPI/reolink-go-api/internal/pkg/network"
)

type RecordingMixin struct{}

type encoding struct {
	audio         bool
	mainBitRate   int
	mainFrameRate int
	mainProfile   string
	mainSize      string
	subBitRate    int
	subFrameRate  int
	subProfile    string
	subSize       string
}

type OptionRecordingEncoding func(*encoding)

// Get the camera's current encoding settings for "Clear" and "Fluent" profiles
// See examples/response/GetEnc.json for example response data
func (rm *RecordingMixin) GetRecordingEncoding() func(handler *network.RestHandler) (*models.Encoding, error) {
	return func(handler *network.RestHandler) (*models.Encoding, error) {
		payload := map[string]interface{}{
			"cmd":    "GetEnc",
			"action": 1,
			"param": map[string]interface{}{
				"channel": 0,
			},
		}

		result, err := handler.Request("POST", payload, "GetEnc", true)

		if err != nil {
			return nil, err
		}

		var recordingData *models.Encoding

		err = json.Unmarshal(result.Value["Enc"], &recordingData)

		if err != nil {
			return nil, err
		}

		return recordingData, nil
	}
}

// Get the recoding advanced setup data
// See examples/response/GetRec.json for example response data
func (rm *RecordingMixin) GetRecordingAdvanced() func(handler *network.RestHandler) (*models.Recording, error) {
	return func(handler *network.RestHandler) (*models.Recording, error) {
		payload := map[string]interface{}{
			"cmd":    "GetRec",
			"action": 1,
			"param": map[string]interface{}{
				"channel": 0,
			},
		}

		result, err := handler.Request("POST", payload, "GetRec", true)

		if err != nil {
			return nil, err
		}

		var encodingData *models.Recording

		err = json.Unmarshal(result.Value["Rec"], &encodingData)

		if err != nil {
			return nil, err
		}

		return encodingData, err

	}
}

// Set the current camera encoding settings for "Clear" and "Fluent" profiles
// Accepts optional parameters of OptionRecordingEncoding type
// Defaults:
// Audio: false
// MainBitRate: 8192
// MainFrameRate: 8
// MainProfile: "High"
// MainSize: "2560*1440"
// SubBitRate: 160
// SubFrameRate: 7
// SubProfile: "High"
// SubSize: "640*480"
func (rm *RecordingMixin) SetRecordingEncoding(encodingOptions ...OptionRecordingEncoding) func(handler *network.RestHandler) (bool, error) {
	encoding := &encoding{
		audio:         false,
		mainBitRate:   8192,
		mainFrameRate: 8,
		mainProfile:   "High",
		mainSize:      "2560*1440",
		subBitRate:    160,
		subFrameRate:  7,
		subProfile:    "High",
		subSize:       "640*480",
	}

	for _, op := range encodingOptions {
		op(encoding)
	}
	return func(handler *network.RestHandler) (bool, error) {
		payload := map[string]interface{}{
			"cmd":    "SetEnc",
			"action": 0,
			"param": map[string]interface{}{
				"Enc": map[string]interface{}{
					"audio":   encoding.audio,
					"channel": 0,
					"mainStream": map[string]interface{}{
						"bitRate":   encoding.mainBitRate,
						"frameRate": encoding.mainFrameRate,
						"profile":   encoding.mainProfile,
						"size":      encoding.mainSize,
					},
					"subStream": map[string]interface{}{
						"bitRate":   encoding.subBitRate,
						"frameRate": encoding.subFrameRate,
						"profile":   encoding.subProfile,
						"size":      encoding.subSize,
					},
				},
			},
		}

		result, err := handler.Request("POST", payload, "SetEnc", true)

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

		return false, fmt.Errorf("camera could not set encoding(s). camera responded with %v", result.Value)
	}
}

// Set audio on or off
// Default: false
func RecordingEncodingOptionAudio(audio bool) OptionRecordingEncoding {
	return func(encoding *encoding) {
		encoding.audio = audio
	}
}

// Set the main bit rate
// Default: 8192
func RecordingEncodingOptionMainBitRate(bitRate enum.MainBitRate) OptionRecordingEncoding {
	return func(encoding *encoding) {
		encoding.mainBitRate = bitRate.Value()
	}
}

// Set the main frame rate
// Default: 8
func RecordingEncodingOptionMainFrameRate(frameRate enum.MainFrameRate) OptionRecordingEncoding {
	return func(encoding *encoding) {
		encoding.mainFrameRate = frameRate.Value()
	}
}

// Set the main profile
// Default: High
func RecordingEncodingOptionMainProfile(profile enum.RecordingProfile) OptionRecordingEncoding {
	return func(encoding *encoding) {
		encoding.mainProfile = profile.Value()
	}
}

// Set the main size
// Default: 2560*1440
func RecordingEncodingOptionMainSize(size enum.MainSize) OptionRecordingEncoding {
	return func(encoding *encoding) {
		encoding.mainSize = size.Value()
	}
}

// Set the sub bit rate
// Default: 160
func RecordingEncodingOptionSubBitRate(bitRate enum.SubBitRate) OptionRecordingEncoding {
	return func(encoding *encoding) {
		encoding.subBitRate = bitRate.Value()
	}
}

// Set the sub frame rate
// Default: 7
func RecordingEncodingOptionSubFrameRate(frameRate enum.SubFrameRate) OptionRecordingEncoding {
	return func(encoding *encoding) {
		encoding.subFrameRate = frameRate.Value()
	}
}

// Set the sub profile
// Default: High
func RecordingEncodingOptionSubProfile(profile enum.RecordingProfile) OptionRecordingEncoding {
	return func(encoding *encoding) {
		encoding.subProfile = profile.Value()
	}
}

// Set the sub size
// Default: 640*480
func RecordingEncodingOptionSubSize(size enum.SubSize) OptionRecordingEncoding {
	return func(encoding *encoding) {
		encoding.subSize = size.Value()
	}
}