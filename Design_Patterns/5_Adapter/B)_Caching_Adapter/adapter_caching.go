package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"strings"
)

/*
	Let's take the previous example.
	In order to actually draw lines as pixels, we turned every single line into points.
	These points are stored in the AddLineCached function.
	This becomes a big of a problem for storage when the function is called repeatedly.
	So, we are going to change the AddLineCached function so we can implement a cache. This cache will prevent us
	from making unnecessary extra data.
*/

type Line struct {
	X1, Y1, X2, Y2 int
}

type VectorImage struct {
	Lines []Line
}

func NewRectangle(width, height int) *VectorImage {
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

type Point struct {
	X, Y int
}

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

// Cache map
var pointCache = map[[16]byte][]Point{}

func (a *vectorToRasterAdapter) AddLineCached(line Line) {
	// Before adding a line, we are going to calculate the lines hash.
	hash := func(obj interface{}) [16]byte { // returns a md5 hash
		bytes, _ := json.Marshal(obj)
		return md5.Sum(bytes)
	}
	h := hash(line)
	// If this line is already on our map, add them to the adapter
	if pts, ok := pointCache[h]; ok {
		for _, pt := range pts {
			a.points = append(a.points, pt)
		}
		return
	}

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

	// Adding a new line to the cache
	pointCache[h] = a.points

	fmt.Println("we have", len(a.points), "points")
}

/*
	This implementation could be improved even further by not storig those extra points, but instead using
	point pointers.

*/

func (v vectorToRasterAdapter) GetPoints() []Point {
	return v.points
}

func VectorToRaster(vi *VectorImage) RasterImage {
	adapter := vectorToRasterAdapter{}

	for _, line := range vi.Lines {
		adapter.AddLineCached(line)
	}

	return adapter
}

func main() {
	rc := NewRectangle(6, 4)
	a := VectorToRaster(rc)

	fmt.Println(DrawPoints(a))
}
