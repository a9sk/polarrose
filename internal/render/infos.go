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
	// h, w, err := terminal.GetFullSize()
	// if err != nil { return }

	if err := terminal.MovCursor(row, col); err != nil {
		panic(fmt.Errorf("%s", err))
	}

	lines := models.GetLines(*infos)

	// print sysinfo in a formatted way, using color if desired
	// fmt.Print(models.ColorCodes[models.CurrentColor])
	fmt.Print(models.ColorCodes["cyan"])

	for i, line := range lines {
		if err := terminal.MovCursor(row+i, col); err != nil {
			panic(fmt.Errorf("%s", err))
		}
		fmt.Printf("%s\n", line)
	}

	fmt.Print(models.ColorCodes["reset"])

}
