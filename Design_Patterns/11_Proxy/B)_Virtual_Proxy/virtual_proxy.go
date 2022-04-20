package main

import "fmt"

/*
	A virtual proxy is a kind of proxy that pretends it is really there when it is not neccesarily.

	Let's imagine that we have an interface called image.
*/

type Image interface {
	Draw()
}

type Bitmap struct {
	filename string
}

func NewBitmap(filename string) *Bitmap {
	fmt.Println("Loading image from ", filename)
	return &Bitmap{filename: filename}
}

func DrawImage(image Image) {
	fmt.Println("About to draw the image")
	image.Draw()
	fmt.Println("Done drawing the image")
}

func (b *Bitmap) Draw() {
	fmt.Println("Drawing image ", b.filename)
}

func main() {
	bmp := NewBitmap("demo.png")
	DrawImage(bmp)
}

/*
	What is the problem with this scenerario?
	The problem is what happens if we never draw the image in the first place, if we never
	invoke DrawImage()

	func main() {
		_ = NewBitmap("demo.png")
		// DrawImage(bmp)
	}

	If we run this, we'll see that there is a fairly obvious problem, and that is that
	we are still loading the image even though we never draw it.
	One attempt to fix this might be to introduce some kind of lazy bitmap,
	the kind where the image doesn't get loaded until we actually need to render it.
	How would be implement this? Using a proxy.

*/

type LazyBitmap struct {
	filename string
	bitmap   *Bitmap
}

func NewLazyBitmap(filename string) *LazyBitmap {
	return &LazyBitmap{filename: filename}
}

/*
	When we make this constructor, we don't initialize the bitmap yet because it is going
	to be lazily constructed.
*/

func (l *LazyBitmap) Draw() {
	if l.bitmap == nil {
		l.bitmap = NewBitmap(l.filename)
	}
	l.bitmap.Draw()
}

func main2() {
	bmp := NewLazyBitmap("demo.png")
	DrawImage(bmp)
	DrawImage(bmp)
	// The second time, the "Loading image from..." message will not be shown.
}

/*
	The demonstration here shows how we can build something tipically called a virtual proxy.
	The reason it is virtual is because when we create a lazy bitmap using the new lazy bitmap function,
	it hasn't been materialized yet, meaning that the underlying implementation of the bitmap
	has not even been constructed. It is only constructed whenever someone explicitly asks for it
*/
