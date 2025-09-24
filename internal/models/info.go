package models

type Info struct {
	OS       string
	Arch     string
	Kernel   string
	Version  string
	Uptime   string
	Hostname string
	Platform string
	CPU      string
	GPU      string
	Memory   string
}

var nInfo = 10 //! remember to update if adding more info fields

func GetNInfo() int {
	return nInfo
}
