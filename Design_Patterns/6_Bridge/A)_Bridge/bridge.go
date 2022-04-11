package main

import "fmt"

/*
	Imagine we are working on a graphical application.
	This application should be capable ofd printing different objects like circles, rectangles, squares, etc.
	However, it need to be able to render them in different ways, like vector form or raster form.
	So, we risk to end up implementing this

							render
							/		\
						circle		square
						/	\			/	\
					raster	vector	raster	vector

	Let's avoid this approach usign the bridge design pattern
*/

type Renderer interface {
	RenderCircle(radius float32)
}

type VectorRenderer struct {
	//...
}

func (v *VectorRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing a circle of radius", radius)
}

type RasterRenderer struct {
	Dpi int
}

func (r *RasterRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing pixels for circle of radius", radius)
}

// We have a single Circle struct that has a bridge to a renderer
type Circle struct {
	renderer Renderer
	radius   float32
}

func NewCircle(renderer Renderer, radius float32) *Circle {
	return &Circle{renderer: renderer, radius: radius}
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

func (c *Circle) Resize(factor float32) {
	c.radius *= factor
}

func main() {
	// Instantiating renderers to introduce as a circle dependency
	raster := RasterRenderer{}
	vector := VectorRenderer{}

	circle := NewCircle(&raster, 5)
	circle2 := NewCircle(&vector, 8)

	circle.Draw()
	circle2.Resize(.5)
}
