package cmd

import (
	"github.com/a9sk/polarrose/internal/render"
	"github.com/a9sk/polarrose/internal/rose"
)

// TODO: migrate width and height to terminal size detection
const (
	width  = 40
	height = 20
	a      = 10.0 // amplitude (radius) of the rose
	k      = 7    // number of petals
	steps  = 100000
	char   = ':' // temprorary character for drawing
)

func Root() {

	points := rose.GenerateRosePoints(a, k, steps)

	render.DrawASCII(points, width, height, char, a)
}
