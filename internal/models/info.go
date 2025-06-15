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

// TODO: implement priority levels for Info
