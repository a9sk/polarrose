package sysinfo

import (
	"fmt"

	"github.com/a9sk/polarrose/internal/models"
	"github.com/shirou/gopsutil/v4/mem"
)

// the struct below is used to gather memory information
// REFERENCE: https://pkg.go.dev/github.com/shirou/gopsutil/v4/mem
//
//
// type VirtualMemoryStat struct {
//
//  Total amount of RAM on this system
// 	Total uint64 `json:"total"`
//
// 	RAM available for programs to allocate (this value is computed from the kernel specific values):
// 	Available uint64 `json:"available"`
//
// 	RAM used by programs (this value is computed from the kernel specific values):
// 	Used uint64 `json:"used"`
//
//  Percentage of RAM used by programs (this value is computed from the kernel specific values):
// 	UsedPercent float64 `json:"usedPercent"`
//
// 	This is the kernel's notion of free memory; RAM chips whose bits nobody
// 	cares about the value of right now. For a human consumable number,
// 	Available is what you really want.
// 	Free uint64 `json:"free"`
//
// 	OS X / BSD specific numbers (http://www.macyourself.com/2010/02/17/what-is-free-wired-active-and-inactive-system-memory-ram/):
// 	Active   uint64 `json:"active"`
// 	Inactive uint64 `json:"inactive"`
// 	Wired    uint64 `json:"wired"`
//
// 	FreeBSD specific numbers (https://reviews.freebsd.org/D8467):
// 	Laundry uint64 `json:"laundry"`
//
// 	Linux specific numbers
// 	https://www.centos.org/docs/5/html/5.1/Deployment_Guide/s2-proc-meminfo.html
// 	https://www.kernel.org/doc/Documentation/filesystems/proc.txt
// 	https://www.kernel.org/doc/Documentation/vm/overcommit-accounting
// 	https://www.kernel.org/doc/Documentation/vm/transhuge.txt
//
//  Buffers        uint64 `json:"buffers"`
// 	Cached         uint64 `json:"cached"`
// 	WriteBack      uint64 `json:"writeBack"`
// 	Dirty          uint64 `json:"dirty"`
// 	WriteBackTmp   uint64 `json:"writeBackTmp"`
// 	Shared         uint64 `json:"shared"`
// 	Slab           uint64 `json:"slab"`
// 	Sreclaimable   uint64 `json:"sreclaimable"`
// 	Sunreclaim     uint64 `json:"sunreclaim"`
// 	PageTables     uint64 `json:"pageTables"`
// 	SwapCached     uint64 `json:"swapCached"`
// 	CommitLimit    uint64 `json:"commitLimit"`
// 	CommittedAS    uint64 `json:"committedAS"`
// 	HighTotal      uint64 `json:"highTotal"`
// 	HighFree       uint64 `json:"highFree"`
// 	LowTotal       uint64 `json:"lowTotal"`
// 	LowFree        uint64 `json:"lowFree"`
// 	SwapTotal      uint64 `json:"swapTotal"`
// 	SwapFree       uint64 `json:"swapFree"`
// 	Mapped         uint64 `json:"mapped"`
// 	VmallocTotal   uint64 `json:"vmallocTotal"`
// 	VmallocUsed    uint64 `json:"vmallocUsed"`
// 	VmallocChunk   uint64 `json:"vmallocChunk"`
// 	HugePagesTotal uint64 `json:"hugePagesTotal"`
// 	HugePagesFree  uint64 `json:"hugePagesFree"`
// 	HugePagesRsvd  uint64 `json:"hugePagesRsvd"`
// 	HugePagesSurp  uint64 `json:"hugePagesSurp"`
// 	HugePageSize   uint64 `json:"hugePageSize"`
// 	AnonHugePages  uint64 `json:"anonHugePages"`
// }

func getMemoryInfo(info *models.Info) error {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return err
	}

	// total memory in GB.
	totalGB := float64(memInfo.Total) / (1024 * 1024 * 1024)
	info.Memory = fmt.Sprintf("%.1f GB", totalGB)

	// used memory in GB.
	usedGB := float64(memInfo.Used) / (1024 * 1024 * 1024)
	info.MemoryUsed = fmt.Sprintf("%.1f GB (%.1f%%)", usedGB, memInfo.UsedPercent)

	// available memory in GB.
	availableGB := float64(memInfo.Available) / (1024 * 1024 * 1024)
	info.MemoryFree = fmt.Sprintf("%.1f GB", availableGB)

	return nil
}
