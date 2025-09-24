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

	// TODO: count the cores
	if len(cpuInfo) > 0 {
		info.CPU = fmt.Sprintf("%d", len(cpuInfo))
	} else {
		info.CPU = ""
	}

	// TODO: retrive and parse more CPUs related informations

	return nil
}
