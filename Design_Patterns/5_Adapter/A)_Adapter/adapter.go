package main

import (
	"fmt"
	"strings"
)

/*
	Let's imagine we are working with an API for rendering graphical objects.
	That API is completely vector base. Everthing is built up out of lines.
*/

type Line struct {
	X1, Y1, X2, Y2 int
}

type VectorImage struct {
	Lines []Line
}

func NewRectangle(width, height int) *VectorImage {
	/*
		The reason why there is a -1 for the width and height is that typically
		arrays and other structures are zero based when you start them out.
		So, if you want an image that has a width of five, it has to go from position 0 to 4.
	*/
	width -= 1
	height -= 1
	return &VectorImage{
		[]Line{
			{0, 0, width, 0},
			{0, 0, 0, height},
			{width, 0, width, height},
			{0, height, width, height},
		},
	}
}

// ↑↑↑↑↑↑↑↑ This is the interface you are given. ↑↑↑↑↑↑↑↑

/*
	The interface we have is going to deal strictly in terms of points, not in terms of lines.
*/

type Point struct {
	X, Y int
}

// A Raster image is an image defined by points, or pixels on the screen
type RasterImage interface {
	GetPoints() []Point
}

func DrawPoints(owner RasterImage) string {
	maxX, maxY := 0, 0
	points := owner.GetPoints()
	for _, pixel := range points {
		if pixel.X > maxX {
			maxX = pixel.X
		}
		if pixel.Y > maxY {
			maxY = pixel.Y
		}
	}

	maxX += 1
	maxY += 1

	data := make([][]rune, maxY)
	for i := 0; i < maxY; i++ {
		data[i] = make([]rune, maxX)
		for j := range data[i] {
			data[i][j] = ' '
		}
	}

	for _, point := range points {
		data[point.Y][point.X] = '*'
	}

	b := strings.Builder{}
	for _, line := range data {
		b.WriteString(string(line))
		b.WriteRune('\n')
	}

	return b.String()
}

// We need an adapter to take the vector image and convert it to a point image
type vectorToRasterAdapter struct {
	points []Point
}

func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}

// This is the adapter
func (a *vectorToRasterAdapter) AddLine(line Line) {
	// Here we need to decompose a line and set up a bunch of points
	left, right := minmax(line.X1, line.X2)
	top, bottom := minmax(line.Y1, line.Y2)
	dx := right - left
	dy := line.Y2 - line.Y1

	if dx == 0 {
		for y := top; y <= bottom; y++ {
			a.points = append(a.points, Point{left, y})
		}
	} else if dy == 0 {
		for x := left; x <= right; x++ {
			a.points = append(a.points, Point{x, top})
		}
	}
}

func (v vectorToRasterAdapter) GetPoints() []Point {
	return v.points
}

func VectorToRaster(vi *VectorImage) RasterImage {
	adapter := vectorToRasterAdapter{}

	for _, line := range vi.Lines {
		adapter.AddLine(line)
	}

	return adapter
}

func main() {
	rc := NewRectangle(6, 4)
	a := VectorToRaster(rc)

	fmt.Println(DrawPoints(a))
}
