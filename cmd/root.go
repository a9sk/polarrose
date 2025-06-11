package cmd

import (
	"github.com/a9sk/polarrose/internal/render"
	"github.com/a9sk/polarrose/internal/rose"
)

// TODO: migrate width and height to terminal size detection
const (
	width  = 40
	height = 20
	steps  = 100000
	char   = ':' // temprorary character for drawing
)

func Root(size float64, petals int) {
	// TODO: validate input, comes from main for now so no need to do that (validated there)

	points := rose.GenerateRosePoints(size, petals, steps)

	render.DrawASCII(points, width, height, char, size)
}
