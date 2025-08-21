package rose

import (
	"fmt"
	"math"

	"github.com/a9sk/polarrose/internal/models"
	"github.com/a9sk/polarrose/internal/terminal"
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

// holds the necessary parameters for mapping between float coordinates and grid coordinates.
type gridConfig struct {
	minX, maxX, minY, maxY float64
	width, height          int
	scaleX, scaleY         float64
}

// creates a new grid configuration based on the external points and padding.
func newGridConfig(externalPoints []models.Point, padding float64) (*gridConfig, error) {
	if len(externalPoints) == 0 {
		return nil, fmt.Errorf("externalPoints cannot be empty")
	}

	// get terminal-calculated rose display size
	terminalWidth, terminalHeight, err := terminal.GetRoseSize()
	if err != nil {
		return nil, fmt.Errorf("failed to get terminal rose size for grid config: %w", err)
	}

	// compute Bounding Box
	minX, maxX := externalPoints[0].X, externalPoints[0].X
	minY, maxY := externalPoints[0].Y, externalPoints[0].Y
	for _, p := range externalPoints {
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	// add padding to the bounding box
	minX -= padding
	maxX += padding
	minY -= padding
	maxY += padding

	// use terminal-calculated dimensions for the grid
	width := terminalWidth
	height := terminalHeight

	// scale factors for mapping float coordinates to grid integers
	// ensure maxX-minX and maxY-minY are not zero to prevent division by zero
	rangeX := maxX - minX
	rangeY := maxY - minY

	if rangeX == 0 || rangeY == 0 {
		return nil, fmt.Errorf("bounding box has zero range, cannot scale")
	}

	scaleX := float64(width-1) / rangeX
	scaleY := float64(height-1) / rangeY

	return &gridConfig{
		minX:   minX,
		maxX:   maxX,
		minY:   minY,
		maxY:   maxY,
		width:  width,
		height: height,
		scaleX: scaleX,
		scaleY: scaleY,
	}, nil
}

// convert a float point to grid coordinates.
func (gc *gridConfig) toGridCoords(p models.Point) (int, int) {
	gx := int(math.Round((p.X - gc.minX) * gc.scaleX))
	gy := int(math.Round((p.Y - gc.minY) * gc.scaleY))
	// we need to ensure coordinates are within grid bounds
	if gx < 0 {
		gx = 0
	} else if gx >= gc.width {
		gx = gc.width - 1
	}
	if gy < 0 {
		gy = 0
	} else if gy >= gc.height {
		gy = gc.height - 1
	}
	return gx, gy
}

// convertt grid coordinates back to float points.
func (gc *gridConfig) fromGridCoords(gx, gy int) models.Point {
	x := float64(gx)/gc.scaleX + gc.minX
	y := float64(gy)/gc.scaleY + gc.minY
	return models.Point{X: x, Y: y}
}

// uses a simple DDA-like algorithm to mark points on the grid between two given points.
func drawLineOnGrid(grid [][]bool, gc *gridConfig, p1, p2 models.Point) {
	x1, y1 := gc.toGridCoords(p1)
	x2, y2 := gc.toGridCoords(p2)

	dx := x2 - x1
	dy := y2 - y1
	steps := int(math.Abs(float64(dx)))
	if int(math.Abs(float64(dy))) > steps {
		steps = int(math.Abs(float64(dy)))
	}

	if steps == 0 { // p1 and p2 are the same grid point
		if x1 >= 0 && x1 < gc.width && y1 >= 0 && y1 < gc.height {
			grid[y1][x1] = true
		}
		return
	}

	xIncrement := float64(dx) / float64(steps)
	yIncrement := float64(dy) / float64(steps)

	for s := 0; s <= steps; s++ {
		currX := int(math.Round(float64(x1) + float64(s)*xIncrement))
		currY := int(math.Round(float64(y1) + float64(s)*yIncrement))
		if currX >= 0 && currX < gc.width && currY >= 0 && currY < gc.height {
			grid[currY][currX] = true
		}
	}
}

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
