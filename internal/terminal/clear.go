package terminal

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func ClearScreen() {
	// i need to try multiple clearing methods for better compatibility
	switch runtime.GOOS {
	case "linux", "darwin":
		// for Unix-like systems i try to use 'clear' command as fallback
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			// if 'clear' fails, i try 'tput clear'
			cmd = exec.Command("tput", "clear")
			cmd.Stdout = os.Stdout
			if err := cmd.Run(); err != nil {
				clearScreenANSI()
			}
		}
	case "windows":
		// windows-specific clearing
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			clearScreenANSI()
		}
	default:
		clearScreenANSI()
	}
}

func clearScreenANSI() {
	// try various ANSI codes to clear the screen TODO: check if this does not run into problems.
	fmt.Print("\033[2J\033[3J\033[H\033[0m")
	fmt.Print("\033[2J\033[3J\033[H")
	fmt.Print("\033c")
	fmt.Print("\033[H")
	fmt.Print("\033[22;2t")
}
