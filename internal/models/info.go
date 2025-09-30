package models

import (
	"fmt"
)

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

func GetLines(infos Info) []string {
	// TODO: add all of the info entries to the GetLines
	lines := []string{
		"System Info:",
		fmt.Sprintf("OS:        %s", infos.OS),
		fmt.Sprintf("Arch:      %s", infos.Arch),
		fmt.Sprintf("Kernel:    %s", infos.Kernel),
		fmt.Sprintf("Version:   %s", infos.Version),
		fmt.Sprintf("Uptime:    %s", infos.Uptime),
		fmt.Sprintf("Hostname:  %s", infos.Hostname),
		fmt.Sprintf("Platform:  %s", infos.Platform),
		fmt.Sprintf("CPU:       %s", infos.CPU),
		fmt.Sprintf("CPU Model: %s", infos.CPUModel),
		fmt.Sprintf("Memory:    %s", infos.Memory),
		fmt.Sprintf("Used:      %s", infos.MemoryUsed),
		fmt.Sprintf("Storage:   %s", infos.Storage),
		fmt.Sprintf("GPU:       %s", infos.GPU),
	}

	return lines
}
