package sysinfo

import (
	"fmt"

	"github.com/a9sk/polarrose/internal/models"
	"github.com/shirou/gopsutil/cpu"
)

// the struct below is used to gather host information
// REFERENCE: https://pkg.go.dev/github.com/shirou/gopsutil/v4
//
//
// type InfoStat struct {
// 	CPU        int32    `json:"cpu"`
// 	VendorID   string   `json:"vendorId"`
// 	Family     string   `json:"family"`
// 	Model      string   `json:"model"`
// 	Stepping   int32    `json:"stepping"`
// 	PhysicalID string   `json:"physicalId"`
// 	CoreID     string   `json:"coreId"`
// 	Cores      int32    `json:"cores"`
// 	ModelName  string   `json:"modelName"`
// 	Mhz        float64  `json:"mhz"`
// 	CacheSize  int32    `json:"cacheSize"`
// 	Flags      []string `json:"flags"`
// 	Microcode  string   `json:"microcode"`
// }

func getCPUInfo(info *models.Info) error {

	cpuInfo, err := cpu.Info()
	if err != nil {
		return err
	}

	if len(cpuInfo) > 0 {
		// CPU count (cores/threads)
		// HACK: i am doing cores and threads in one single thing, TODO: split it in two different parts
		if logical, err := cpu.Counts(true); err == nil {
			if physical, err := cpu.Counts(false); err == nil {
				info.CPU = fmt.Sprintf("%d cores, %d threads", physical, logical)
			}
		} else {
			info.CPU = "Unknown"
		}

		// CPU model name
		info.CPUModel = cpuInfo[0].ModelName

		// Physical cores
		// info.CPUCores = strconv.Itoa(int(cpuInfo[0].Cores))

		// Logical processors (threads)
		// logicalCount, err := cpu.Counts(true)
		// if err == nil {
		// 	info.CPUThreads = strconv.Itoa(logicalCount)
		// }
	} else {
		info.CPU = "Unknown"
		info.CPUModel = "Unknown"
		info.CPUCores = "Unknown"
		info.CPUThreads = "Unknown"
	}

	return nil
}
