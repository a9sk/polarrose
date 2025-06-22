package rose

import (
	"math"

	"github.com/a9sk/polarrose/internal/models"
)

// GenerateRosePoints generates points for a rose curve in polar coordinates.
func GenerateRosePoints(a float64, k int, steps int) []models.Point {
	points := make([]models.Point, 0, steps)
	for i := 0; i < steps; i++ {
		theta := 2 * math.Pi * float64(i) / float64(steps) // theta is the angle between 0 and 2π
		// apply the rose formula: r = a * cos(k * theta)
		r := a * math.Cos(float64(k)*theta)
		x := r * math.Cos(theta)
		y := r * math.Sin(theta)

		points = append(points, models.Point{X: x, Y: y})
	}
	return points
}

// TODO: add a function that "fills" the points with ascii characters
