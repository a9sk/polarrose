package render

import (
	"fmt"

	"github.com/a9sk/polarrose/internal/models"
)

func DrawInfo(infos *models.Info) {
	// TODO: write the sysinfo next to the rose using the terminal.GetInfoPos() function
	// TODO: when writing the info, check the terminal size with terminal.GetFullSize()

	fmt.Print(infos)
}
