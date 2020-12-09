package models

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

type DevicePerformanceInformation struct {
	CodecRate     int `json:"codecRate"`
	CpuUsed       int `json:"cpuUsed"`
	NetThroughput int `json:"netThroughput"`
}

// TODO: update
type DeviceReboot struct{}
