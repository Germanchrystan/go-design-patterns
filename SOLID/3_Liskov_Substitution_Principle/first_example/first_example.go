package firstexample

import "fmt"

/*
Let's suppose you are trying to deal with geometric shapes of a rectangular nature.
*/

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Print("Expected an area of ", expectedArea, ", but got ", actualArea, "\n")
}

// Now lets create a Square struct from a composite of Rectangle
type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

// This is where the principle is broken.
// In order to keep the square shape, the setters for this struct must keep both ints equal
func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(height int) {
	s.height = height
	s.width = height
}

/*
If you are expecting some sort of behaviour up the hierarchy, it should continue to work,
even if you proceed to extend objects and execute the same methods with them.
*/
func main() {
	rc := &Rectangle{2, 3}
	UseIt(rc) // "Expected an areo of 20, but got 20"

	sq := NewSquare(5)
	UseIt(sq) // "Expected an areo of 50, but got 100"
	/*
		This is the result of the setters working differently for the Square.
		The LSP states that if you continue to use generalizations (like interfaces),
		then you should not have inherited or you should not have implementations of those generalizations
		that break some of the assumptions which are set up at the higher level.

		The behaviour of implementors of a particular type should not break the core fundamental
		behaviours that you rely on.

	*/
}
