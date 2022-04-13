package main

import "fmt"

/*
	We have geometric shapes in the system, and we want to extend their functionality
	by giving them additional properties.
*/

type Shape interface {
	Render() string
}

type Circle struct {
	Radius float32
}

func (c *Circle) Render() string {
	return fmt.Sprintf("Circle of radius %f", c.Radius)
}

func (c *Circle) Resize(factor float32) {
	c.Radius *= factor
}

type Square struct {
	Side float32
}

func (s *Square) Render() string {
	return fmt.Sprintf("Square with side %f", s.Side)
}

/*
	Now, imagine we want to add colors to the shapes.
	We could add an additional property to the structs,
	but we would break the Open Closed Principle.

	We should extend these types.
*/

type ColoredSquare struct {
	Square
	Color string
}

/*
	This could work if you have a small number of structs to extend.
	If there are a lot of different shape structs,
	having a counterpart colored struct of every single one is not scalable.
*/
type ColoredShape struct {
	Shape Shape
	Color string
}

func (c *ColoredShape) Render() string {
	return fmt.Sprintf("%s has the color %s", c.Shape.Render(), c.Color)
}

func main() {
	circle := Circle{2}
	fmt.Println(circle.Render())

	redCircle := ColoredShape{&circle, "Red"}
	fmt.Println(redCircle.Render())
}

/*
	There are certain thing we lose with a decorator.
	In this example above, circle struct could be resized, like this:

	circle.Resize(2)

	However, the resize method can't be applied to the ColoredShape struct

	redCircle.Resize(2) => Error

	It is only the Circle struct that has the resize method.

	The upside is that decorators can be composed, which means we can apply decorators to decorators.
*/

type TransparentShape struct {
	Shape        Shape
	Transparency float32
}

func (t *TransparentShape) Render() string {
	return fmt.Sprintf("%s has %f%% transparency", t.Shape.Render(), t.Transparency*100.0)
}

func main2() {
	circle := Circle{2}

	redCircle := ColoredShape{&circle, "Red"}
	rhsCircle := TransparentShape{&redCircle, 0.5}
	fmt.Println(rhsCircle.Render())
}

/*
	Decorators can be composed so we can apply decorators to other decorators.
	This does not do any kind of detection in terms of circular dependencies or in terms of repetition.

	You can apply a color shape to a color shape in this example.
	That is not necessarily a problem. If we wanted to start detecting repetitions of decorators,
	then it is going to be a lot more work.

	In some situations, very rare ones, it might be worth detecting duplicated application of decorators,
	but in most cases it is ok.
*/
