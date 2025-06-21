package sysinfo

import (
	"fmt"

	"github.com/a9sk/polarrose/internal/models"
	"github.com/shirou/gopsutil/v4/host"
)

// the struct below is used to gather host information
// REFERENCE: https://pkg.go.dev/github.com/shirou/gopsutil/v4
//
//
// type InfoStat struct {
// 	Hostname             string `json:"hostname"`
// 	Uptime               uint64 `json:"uptime"`
// 	BootTime             uint64 `json:"bootTime"`
// 	Procs                uint64 `json:"procs"`           // number of processes
// 	OS                   string `json:"os"`              // ex: freebsd, linux
// 	Platform             string `json:"platform"`        // ex: ubuntu, linuxmint
// 	PlatformFamily       string `json:"platformFamily"`  // ex: debian, rhel
// 	PlatformVersion      string `json:"platformVersion"` // version of the complete OS
// 	KernelVersion        string `json:"kernelVersion"`   // version of the OS kernel (if available)
// 	KernelArch           string `json:"kernelArch"`      // native cpu architecture queried at runtime, as returned by `uname -m` or empty string in case of error
// 	VirtualizationSystem string `json:"virtualizationSystem"`
// 	VirtualizationRole   string `json:"virtualizationRole"` // guest or host
// 	HostID               string `json:"hostId"`             // ex: uuid
// }

func getHostInfo(info *models.Info) error {

	hostInfo, err := host.Info()
	if err != nil {
		return err
	}

	info.Hostname = hostInfo.Hostname

	info.Uptime = fmt.Sprintf("%d", hostInfo.Uptime) // parse to string

	info.OS = hostInfo.OS

	info.Platform = hostInfo.Platform

	info.Kernel = hostInfo.KernelVersion

	info.Arch = hostInfo.KernelArch

	info.Version = hostInfo.PlatformVersion

	return nil
}
