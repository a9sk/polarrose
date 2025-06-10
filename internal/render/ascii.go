package render

import (
	"fmt"

	"github.com/a9sk/polarrose/internal/models"
)

// TODO: migrate width and height to terminal size detection (considering sysinfo)
func DrawASCII(points []models.Point, w, h int, char rune, a float64) {

	canvas := make([][]rune, h)
	for i := range canvas {
		canvas[i] = make([]rune, w)
		for j := range canvas[i] {
			canvas[i][j] = ' '
		}
	}

	// normalize and map points to terminal grid
	for _, p := range points {
		col := int((p.X/a + 1) * float64(w) / 2)
		row := int((1 - p.Y/a) * float64(h) / 2)
		if row >= 0 && row < h && col >= 0 && col < w {
			canvas[row][col] = char
		}
	}

	for _, row := range canvas {
		fmt.Println(string(row))
	}
}
