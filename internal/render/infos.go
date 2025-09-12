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

	// TODO: move this to internal/terminal function (ANSI escape code: \033[row;colH)
	fmt.Printf("\033[%d;%dH", row, col)

	// print sysinfo in a formatted way, using color if desired
	// fmt.Print(models.ColorCodes[models.CurrentColor])
	fmt.Print(models.ColorCodes["cyan"])
	fmt.Printf("System Info:\n")
	fmt.Printf("\033[%d;%dH", row+1, col)
	fmt.Printf("OS:       %s\n", infos.OS)
	fmt.Printf("\033[%d;%dH", row+2, col)
	fmt.Printf("Arch:     %s\n", infos.Arch)
	fmt.Printf("\033[%d;%dH", row+3, col)
	fmt.Printf("Kernel:   %s\n", infos.Kernel)
	fmt.Printf("\033[%d;%dH", row+4, col)
	fmt.Printf("Version:  %s\n", infos.Version)
	fmt.Printf("\033[%d;%dH", row+5, col)
	fmt.Printf("Uptime:   %s\n", infos.Uptime)
	fmt.Printf("\033[%d;%dH", row+6, col)
	fmt.Printf("Hostname: %s\n", infos.Hostname)
	fmt.Printf("\033[%d;%dH", row+7, col)
	fmt.Printf("Platform: %s\n", infos.Platform)
	fmt.Printf("\033[%d;%dH", row+8, col)
	fmt.Printf("CPU:      %s\n", infos.CPU)
	fmt.Printf("\033[%d;%dH", row+9, col)
	fmt.Printf("GPU:      %s\n", infos.GPU)
	fmt.Printf("\033[%d;%dH", row+10, col)
	fmt.Printf("Memory:   %s\n", infos.Memory)
	fmt.Print(models.ColorCodes["reset"])
}
