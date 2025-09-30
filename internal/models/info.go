package models

type Info struct {
	OS         string
	Arch       string
	Kernel     string
	Version    string
	Uptime     string
	Hostname   string
	Platform   string
	CPU        string
	CPUModel   string
	CPUCores   string
	CPUThreads string
	GPU        string
	Memory     string
	MemoryUsed string
	MemoryFree string
	Storage    string
}

var nInfo = 16

func GetNInfo() int {
	return nInfo
}
