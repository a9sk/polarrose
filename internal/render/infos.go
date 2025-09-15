package render

import (
	"fmt"

	"github.com/a9sk/polarrose/internal/models"
	"github.com/a9sk/polarrose/internal/terminal"
)

func DrawInfo(infos *models.Info) {

	// fmt.Print(infos)

	col, row := terminal.GetInfoPos()

	// optionally, check terminal size (not strictly needed for just printing)
	// w, h, err := terminal.GetFullSize()
	// if err != nil { return }

	if err := terminal.MovCursor(row, col); err != nil {
		panic(fmt.Errorf("%s", err))
	}

	// print sysinfo in a formatted way, using color if desired
	// fmt.Print(models.ColorCodes[models.CurrentColor])
	fmt.Print(models.ColorCodes["cyan"])
	fmt.Printf("System Info:\n")
	if err := terminal.MovCursor(row+1, col); err != nil {
		panic(fmt.Errorf("%s", err))
	}
	fmt.Printf("OS:       %s\n", infos.OS)
	if err := terminal.MovCursor(row+2, col); err != nil {
		panic(fmt.Errorf("%s", err))
	}
	fmt.Printf("Arch:     %s\n", infos.Arch)
	if err := terminal.MovCursor(row+3, col); err != nil {
		panic(fmt.Errorf("%s", err))
	}
	fmt.Printf("Kernel:   %s\n", infos.Kernel)
	if err := terminal.MovCursor(row+4, col); err != nil {
		panic(fmt.Errorf("%s", err))
	}
	fmt.Printf("Version:  %s\n", infos.Version)
	if err := terminal.MovCursor(row+5, col); err != nil {
		panic(fmt.Errorf("%s", err))
	}
	fmt.Printf("Uptime:   %s\n", infos.Uptime)
	if err := terminal.MovCursor(row+6, col); err != nil {
		panic(fmt.Errorf("%s", err))
	}
	fmt.Printf("Hostname: %s\n", infos.Hostname)
	if err := terminal.MovCursor(row+7, col); err != nil {
		panic(fmt.Errorf("%s", err))
	}
	fmt.Printf("Platform: %s\n", infos.Platform)
	if err := terminal.MovCursor(row+8, col); err != nil {
		panic(fmt.Errorf("%s", err))
	}
	fmt.Printf("CPU:      %s\n", infos.CPU)
	if err := terminal.MovCursor(row+9, col); err != nil {
		panic(fmt.Errorf("%s", err))
	}
	fmt.Printf("GPU:      %s\n", infos.GPU)
	if err := terminal.MovCursor(row+10, col); err != nil {
		panic(fmt.Errorf("%s", err))
	}
	fmt.Printf("Memory:   %s\n", infos.Memory)
	fmt.Print(models.ColorCodes["reset"])
}
