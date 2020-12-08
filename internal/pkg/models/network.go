package models

// TODO: update with its actual data structure
type ScanWifi struct {
}

// TODO: update with its actual data structure
type Wifi struct {
}

type NetworkGeneralDns struct {
	Auto int    `json:"auto"`
	Dns1 string `json:"dns1"`
	Dns2 string `json:"dns2"`
}

type NetworkGeneralStatic struct {
	Gateway string `json:"gateway"`
	Ip      string `json:"ip"`
	Mask    string `json:"mask"`
}

type NetworkGeneral struct {
	ActiveLink string               `json:"activeLink"`
	Dns        NetworkGeneralDns    `json:"dns"`
	Mac        string               `json:"mac"`
	Static     NetworkGeneralStatic `json:"static"`
	Type       string               `json:"type"`
}

type NetworkDDNS struct {
	Domain   string `json:"domain"`
	Enable   bool   `json:"enable"`
	Password string `json:"password"`
	Type     string `json:"type"`
	Username string `json:"userName"`
}

type NetworkNTP struct {
	Enable   bool   `json:"enable"`
	Interval int    `json:"interval"`
	Port     int    `json:"port"`
	Server   string `json:"server"`
}

type NetworkSchedule struct {
	Enable bool   `json:"enable"`
	Table  string `json:"table"`
}

type NetworkEmail struct {
	Username   string          `json:"username"`
	Password   string          `json:"password"`
	Addr1      string          `json:"addr1"`
	Addr2      string          `json:"addr2"`
	Addr3      string          `json:"addr3"`
	Attachment string          `json:"attachment"`
	Interval   string          `json:"interval"`
	Nickname   string          `json:"nickName"`
	Schedule   NetworkSchedule `json:"schedule"`
	SmtpPort   int             `json:"smtpPort"`
	SmtpServer string          `json:"smtpServer"`
	SSL        string          `json:"ssl"`
}

type NetworkFTP struct {
	Username   string          `json:"userName"`
	Password   string          `json:"password"`
	Anonymous  bool            `json:"anonymous"`
	Interval   int             `json:"interval"`
	MaxSize    int             `json:"maxSize"`
	Mode       int             `json:"mode"`
	Port       int             `json:"port"`
	RemoteDir  string          `json:"remoteDir"`
	Schedule   NetworkSchedule `json:"schedule"`
	Server     string          `json:"server"`
	StreamType int             `json:"streamType"`
}

type NetworkPush struct {
	Schedule NetworkSchedule `json:"schedule"`
}
