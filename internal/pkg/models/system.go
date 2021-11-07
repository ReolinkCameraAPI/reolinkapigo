package models

import "encoding/json"

type DstInformation struct {
	Enable       bool `json:"enable"`
	EndHour      int  `json:"endHour"`
	EndMin       int  `json:"endMin"`
	EndMon       int  `json:"endMon"`
	EndSec       int  `json:"endSec"`
	EndWeek      int  `json:"endWeek"`
	EndWeekday   int  `json:"endWeekday"`
	Offset       int  `json:"offset"`
	StartHour    int  `json:"startHour"`
	StartMin     int  `json:"startMin"`
	StartMon     int  `json:"startMon"`
	StartSec     int  `json:"startSec"`
	StartWeek    int  `json:"startWeek"`
	StartWeekday int  `json:"startWeekday"`
}

type TimeInformation struct {
	Day      int    `json:"day"`
	Hour     int    `json:"hour"`
	HourFmt  bool   `json:"hourFmt"`
	Min      int    `json:"min"`
	Mon      int    `json:"mon"`
	Sec      int    `json:"sec"`
	TimeFmt  string `json:"timeFmt"`
	TimeZone int    `json:"timeZone"`
	Year     int    `json:"year"`
}

type DeviceInformation struct {
	B485            int    `json:"B485"`
	IoInputNumber   int    `json:"IOInputNum"`
	IoOutputNumber  int    `json:"IOOutputNum"`
	AudioNumber     int    `json:"AudioNum"`
	BuildDay        string `json:"buildDay"`
	ConfigVersion   string `json:"cfgVer"`
	ChannelNumber   int    `json:"channelNum"`
	Detail          string `json:"detail"`
	DiskNumber      int    `json:"diskNum"`
	FirmwareVersion string `json:"firmVer"`
	HardwareVersion string `json:"hardVer"`
	Model           string `json:"model"`
	Name            string `json:"name"`
	Serial          string `json:"serial"`
	Type            string `json:"type"`
	Wifi            bool   `json:"wifi"`
}

func (d *DeviceInformation) UnmarshalJSON(b []byte) error {
	var deviceInformation struct {
		B485            int    `json:"B485"`
		IoInputNumber   int    `json:"IOInputNum"`
		IoOutputNumber  int    `json:"IOOutputNum"`
		AudioNumber     int    `json:"AudioNum"`
		BuildDay        string `json:"buildDay"`
		ConfigVersion   string `json:"cfgVer"`
		ChannelNumber   int    `json:"channelNum"`
		Detail          string `json:"detail"`
		DiskNumber      int    `json:"diskNum"`
		FirmwareVersion string `json:"firmVer"`
		HardwareVersion string `json:"hardVer"`
		Model           string `json:"model"`
		Name            string `json:"name"`
		Serial          string `json:"serial"`
		Type            string `json:"type"`
		Wifi            int    `json:"wifi"`
	}

	if err := json.Unmarshal(b, &deviceInformation); err != nil {
		var deviceInfo struct {
			B485            int    `json:"B485"`
			IoInputNumber   int    `json:"IOInputNum"`
			IoOutputNumber  int    `json:"IOOutputNum"`
			AudioNumber     int    `json:"AudioNum"`
			BuildDay        string `json:"buildDay"`
			ConfigVersion   string `json:"cfgVer"`
			ChannelNumber   int    `json:"channelNum"`
			Detail          string `json:"detail"`
			DiskNumber      int    `json:"diskNum"`
			FirmwareVersion string `json:"firmVer"`
			HardwareVersion string `json:"hardVer"`
			Model           string `json:"model"`
			Name            string `json:"name"`
			Serial          string `json:"serial"`
			Type            string `json:"type"`
			Wifi            bool   `json:"wifi"`
		}
		if err := json.Unmarshal(b, &deviceInfo); err != nil {
			return err
		} else {
			d.B485 = deviceInfo.B485
			d.IoInputNumber = deviceInfo.IoInputNumber
			d.IoOutputNumber = deviceInfo.IoOutputNumber
			d.AudioNumber = deviceInfo.AudioNumber
			d.BuildDay = deviceInfo.BuildDay
			d.ConfigVersion = deviceInfo.ConfigVersion
			d.ChannelNumber = deviceInfo.ChannelNumber
			d.DiskNumber = deviceInfo.DiskNumber
			d.FirmwareVersion = deviceInfo.FirmwareVersion
			d.HardwareVersion = deviceInfo.HardwareVersion
			d.Model = deviceInfo.Model
			d.Name = deviceInfo.Name
			d.Serial = deviceInfo.Serial
			d.Type = deviceInfo.Type
			d.Wifi = deviceInfo.Wifi
			return nil
		}
	}

	switch deviceInformation.Wifi {
	case 1:
		d.Wifi = true
	default:
		d.Wifi = false
	}

	d.B485 = deviceInformation.B485
	d.IoInputNumber = deviceInformation.IoInputNumber
	d.IoOutputNumber = deviceInformation.IoOutputNumber
	d.AudioNumber = deviceInformation.AudioNumber
	d.BuildDay = deviceInformation.BuildDay
	d.ConfigVersion = deviceInformation.ConfigVersion
	d.ChannelNumber = deviceInformation.ChannelNumber
	d.DiskNumber = deviceInformation.DiskNumber
	d.FirmwareVersion = deviceInformation.FirmwareVersion
	d.HardwareVersion = deviceInformation.HardwareVersion
	d.Model = deviceInformation.Model
	d.Name = deviceInformation.Name
	d.Serial = deviceInformation.Serial
	d.Type = deviceInformation.Type

	return nil
}

type DevicePerformanceInformation struct {
	CodecRate     int `json:"codecRate"`
	CpuUsed       int `json:"cpuUsed"`
	NetThroughput int `json:"netThroughput"`
}

type DeviceNorm struct {
	Norm string `json:"norm"`
}

type DeviceGeneralInformation struct {
	Time *TimeInformation
	Dst  *DstInformation
	Norm *DeviceNorm
}
