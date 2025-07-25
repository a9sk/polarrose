package terminal

import "fmt"

func GetInfoPos() (int, int) {
	// this function looks at where the rose "ends" and returns the positions
	// where the sysinfo should be printed

	x, _, err := GetRoseSize()
	if err != nil {
		// this should panic or we would overwrite the rose
		panic(fmt.Sprintf("PANIC: %v", err))
	}

	return x + 2, 1
}
