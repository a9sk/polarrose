package sysinfo

import (
	"fmt"

	"github.com/a9sk/polarrose/internal/models"
	"github.com/shirou/gopsutil/v4/host"
)

// the structs below are used to gather system information
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

// TODO: better error handling not limiting the previously fetched data

func GetSysInfo() (*models.Info, error) {
	// i don't think this is how you do it
	err := error(nil)

	// we should gather system information here and return a models.Info struct
	info := &models.Info{}

	if err = getHostInfo(info); err != nil {
		return nil, fmt.Errorf("failed to get host info: %w", err)
	}

	// TODO: gather more information like CPU, GPU, Memory, etc.

	return info, nil
}

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
