// this package should be used to get the terminal size
// TODO: have listeners for terminal size changes

package terminal

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func GetFullSize() (int, int, error) {

	// TODO: check why powershell does not work
	setTerminalRaw()

	w, h, err := term.GetSize(0) // standard input (stdin) sizes
	if err != nil {
		// maybe a panic might be good here, similar to GetRoseSize
		// (or maybe it is bad there)
		return 0, 0, fmt.Errorf("failed to get terminal size: %w", err)
	}

	// fmt.Printf("terminal size: %d x %d\n", w, h)
	return w, h, nil
}

func GetRoseSize() (int, int, error) {
	w, h, err := GetFullSize()

	if err != nil {
		// panics are pretty neat if i understand correctly
		// here it makes sense because we cannot do anything without a valid terminal size
		// panic(fmt.Sprintf("Error getting terminal size: %v", err))
		return 0, 0, fmt.Errorf("failed to get rose size: %w", err)
	}

	// the width is half the terminal to leave space for sysinfo
	// the height is 85% of the terminal so that it fits nicely
	return w / 2, h - (h * 15 / 100), nil
}

func setTerminalRaw() {

	// this is a workaround to make sure that the terminal is in raw mode
	// it is needed for the term.GetSize to work properly
	// it is not needed for the terminal size detection itself

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)
}
