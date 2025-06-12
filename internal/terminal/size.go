// this package should be used to get the terminal size
// TODO: have listeners for terminal size changes

package terminal

import (
	"fmt"

	"golang.org/x/term"
)

func GetFullSize() (int, int, error) {
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
		panic(fmt.Sprintf("Error getting terminal size: %v", err))
	}

	// the width is half the terminal to leave space for sysinfo
	// TODO: we want to limit the height of the terminal if it is too big
	return w / 2, h, nil
}
