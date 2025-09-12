package terminal

import (
	"fmt"

	"github.com/a9sk/polarrose/internal/models"
)

func GetInfoPos() (int, int) {
	// this function looks at where the rose "ends" and returns the positions
	// where the sysinfo should be printed

	// we want to check for the remaining space's height (which is the full
	// terminal height) and then use it to center the info.
	// we also want to center the info based on the number of lines it will take

	_, row, err := GetFullSize()
	if err != nil {
		panic("PANIC: terminal size is zero")
	}

	x, _, err := GetRoseSize()
	if err != nil {
		// this should panic or we would overwrite the rose
		panic(fmt.Sprintf("PANIC: %v", err))
	}

	tl := models.GetNInfo()
	m := 1

	if row >= tl {
		m = int(row/2) - int(tl/2)
	}

	return x + 5, m
}
