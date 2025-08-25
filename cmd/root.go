package cmd

import (
	"fmt"

	"github.com/a9sk/polarrose/internal/render"
	"github.com/a9sk/polarrose/internal/rose"
	"github.com/a9sk/polarrose/internal/sysinfo"
	"github.com/a9sk/polarrose/internal/terminal"
)

const (
	steps = 100000
	char  = ':' // temprorary character for drawing
)

func Root(size float64, petals int, color string) {
	// TODO: validate input, comes from main for now so no need to do that (validated there)

	width, height, err := terminal.GetRoseSize()
	if err != nil {
		panic(fmt.Sprintf("[FAIL] in TERMINAL: %v", err))
	}

	// show the external points
	points := rose.GenerateRosePoints(size, petals, steps)
	// render.DrawASCII(points, width, height, char, size, color)

	// show the internal points
	pointsIn := rose.GetInternalPoints(points)
	// render.DrawASCII(pointsIn, width, height, char, size, color)

	// show the whole points
	pointsAll := append(points, pointsIn...)
	render.DrawASCII(pointsAll, width, height, char, size, color)

	infos, err := sysinfo.GetSysInfo()
	if err != nil {
		panic(fmt.Sprintf("[FAIL] in SYSINFO: %v", err))
	}

	// show the system information
	render.DrawInfo(infos)
}
