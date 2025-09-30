package sysinfo

import (
	"fmt"

	"github.com/a9sk/polarrose/internal/models"
	"github.com/shirou/gopsutil/v4/disk"
)

// the struct below is used to gather disk/storage information
// REFERENCE: https://pkg.go.dev/github.com/shirou/gopsutil/v4/disk
//
//
// type UsageStat struct {
// 	Path              string  `json:"path"`
// 	Fstype            string  `json:"fstype"`
// 	Total             uint64  `json:"total"`
// 	Free              uint64  `json:"free"`
// 	Used              uint64  `json:"used"`
// 	UsedPercent       float64 `json:"usedPercent"`
// 	InodesTotal       uint64  `json:"inodesTotal"`
// 	InodesUsed        uint64  `json:"inodesUsed"`
// 	InodesFree        uint64  `json:"inodesFree"`
// 	InodesUsedPercent float64 `json:"inodesUsedPercent"`
// }

func getStorageInfo(info *models.Info) error {
	usage, err := disk.Usage("/") // Unix-like systems.
	if err != nil {
		// try Windows root.
		usage, err = disk.Usage("C:")
		if err != nil {
			return err
		}
	}

	totalGB := float64(usage.Total) / (1024 * 1024 * 1024)
	usedGB := float64(usage.Used) / (1024 * 1024 * 1024)
	info.Storage = fmt.Sprintf("%.1f GB / %.1f GB (%.1f%%)", usedGB, totalGB, usage.UsedPercent)

	return nil
}
