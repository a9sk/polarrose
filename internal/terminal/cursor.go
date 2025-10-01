package terminal

import "fmt"

// moves the cursor to any (existing) position
func MovCursor(row, col int) error {

	h, w, err := GetFullSize()
	if err != nil {
		return fmt.Errorf("problem moving the cursor: %s", err)
	}

	if row <= 0 || row >= h || col <= 0 || col >= w {
		return fmt.Errorf("problem moving the cursor: out of bound")
	}

	// as in https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797
	fmt.Printf("\033[%d;%dH", row, col)

	return nil
}
