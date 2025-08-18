package rose

import (
	"fmt"
	"math"

	"github.com/a9sk/polarrose/internal/models"
)

// generates points for a rose curve in polar coordinates.
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

// TODO: add a function that, given a list of points, returns the list of internal points

// calculates and returns the internal points of a rose curve.
func GetInternalPoints(externalPoints []models.Point) []models.Point {
	if len(externalPoints) == 0 {
		return nil
	}

	const padding = 5.0 // adding around the bounding box

	gc, err := newGridConfig(externalPoints, padding)
	if err != nil {
		panic(fmt.Sprintf("error creating grid config: %v\n", err))
	}

	// create a 2D boolean grid to represent the drawing area.
	grid := make([][]bool, gc.height)
	for i := range grid {
		grid[i] = make([]bool, gc.width)
	}

	/*
	 *
	 * I FOLLOW THE STEPS AS DESCRIBED IN THE README.md file in the internal/rose directory.
	 *
	 */

	// 1. boundary drawing (rasterization)
	for i := 0; i < len(externalPoints); i++ {
		p1 := externalPoints[i]
		p2 := externalPoints[(i+1)%len(externalPoints)] // connect last point to first
		drawLineOnGrid(grid, gc, p1, p2)
	}

	// 2. flood fill from all corners (BFS)
	queue := []models.Point{}
	visited := make(map[models.Point]bool) // i use models.Point as key for visited map

	// we need to add the four corners of the grid to the queue (converted to float for consistency with BFS queue type)
	queue = append(queue, gc.fromGridCoords(0, 0))
	queue = append(queue, gc.fromGridCoords(gc.width-1, 0))
	queue = append(queue, gc.fromGridCoords(0, gc.height-1))
	queue = append(queue, gc.fromGridCoords(gc.width-1, gc.height-1))

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		gx, gy := gc.toGridCoords(curr)

		if visited[curr] {
			continue
		}
		visited[curr] = true

		// mark this grid cell as visited (external)
		if gx >= 0 && gx < gc.width && gy >= 0 && gy < gc.height {
			grid[gy][gx] = true // mark as visited "external"
		}

		// explore neighbors (up, down, left, right)
		neighbors := []models.Point{
			gc.fromGridCoords(gx+1, gy),
			gc.fromGridCoords(gx-1, gy),
			gc.fromGridCoords(gx, gy+1),
			gc.fromGridCoords(gx, gy-1),
		}

		for _, neighbor := range neighbors {
			nx, ny := gc.toGridCoords(neighbor)
			// check if neighbor is within bounds and not a boundary point already marked true
			if nx >= 0 && nx < gc.width && ny >= 0 && ny < gc.height && !grid[ny][nx] && !visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
	}

	// 3. collect internal points
	internalPoints := []models.Point{}
	for y := 0; y < gc.height; y++ {
		for x := 0; x < gc.width; x++ {
			if !grid[y][x] { // jf it's not a boundary and not visited by flood fill, it's internal
				internalPoints = append(internalPoints, gc.fromGridCoords(x, y))
			}
		}
	}

	return internalPoints
}
