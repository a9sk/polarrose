package render

import (
	"fmt"

	"github.com/a9sk/polarrose/internal/models"
	"github.com/a9sk/polarrose/internal/terminal"
)

func PrintWatermark() error {

	w, h, err := terminal.GetFullSize()
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	if err := terminal.MovCursor(h-2, w-28); err != nil {
		return fmt.Errorf("%s", err)
	}
	fmt.Print(models.Watermark[0])

	if err := terminal.MovCursor(h-1, w-28); err != nil {
		return fmt.Errorf("%s", err)
	}
	fmt.Print(models.Watermark[1])

	return nil
}
