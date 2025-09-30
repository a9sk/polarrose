package sysinfo

import (
	"fmt"

	"github.com/a9sk/polarrose/internal/models"
)

// TODO: better error handling not limiting the previously fetched data

func GetSysInfo() (*models.Info, error) {
	// i don't think this is how you do it
	err := error(nil)

	// we should gather system information here and return a models.Info struct
	info := &models.Info{}

	if err = getHostInfo(info); err != nil {
		return nil, fmt.Errorf("failed to get [host] info: %w", err)
	}

	if err = getCPUInfo(info); err != nil {
		return nil, fmt.Errorf("failed to get [CPU] info: %w", err)
	}

	if err = getMemoryInfo(info); err != nil {
		return nil, fmt.Errorf("failed to get [memory] info: %w", err)
	}

	if err = getStorageInfo(info); err != nil {
		return nil, fmt.Errorf("failed to get [storage] info: %w", err)
	}

	// TODO: gather more information like CPU, GPU, Memory, etc.

	return info, nil
}
