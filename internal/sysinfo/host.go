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

	info.Uptime = parseUptime(hostInfo.Uptime)

	info.OS = hostInfo.OS

	info.Platform = hostInfo.Platform

	info.Kernel = hostInfo.KernelVersion

	info.Arch = hostInfo.KernelArch

	info.Version = hostInfo.PlatformVersion

	return nil
}

func parseUptime(uptimeSeconds uint64) string {
	// convert uptime in seconds to a human-readable format
	days := uptimeSeconds / 86400
	hours := (uptimeSeconds % 86400) / 3600
	minutes := (uptimeSeconds % 3600) / 60
	seconds := uptimeSeconds % 60

	if days > 0 {
		return fmt.Sprintf("%dd %dh %dm %ds", days, hours, minutes, seconds)
	} else if hours > 0 {
		return fmt.Sprintf("%dh %dm %ds", hours, minutes, seconds)
	} else if minutes > 0 {
		return fmt.Sprintf("%dm %ds", minutes, seconds)
	} else {
		return fmt.Sprintf("%ds", seconds)
	}
}
