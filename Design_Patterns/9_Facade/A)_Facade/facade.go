package main

import "fmt"

/*
	Let's build a multi buffer viewport terminal.
	Sometimes, it is required to have more than one buffer to store the text being outputted.
	So imagine if we were running several programs at the same time.
	We would want to take this entire space and somehow split it into different parts.
	So we have different viewpoints, which in turn are attached to different buffers, but we
	still to work with a console in as much as simplified manner as possible.

	Let's start by making a buffer struct, which stores letters.

*/

type Buffer struct {
	width, height int
	buffer        []rune
}

func NewBuffer(width int, height int) *Buffer {
	return &Buffer{width: width, height: height,
		buffer: make([]rune, width*height)}
}

// Then we could have a utility for getting a character at a particular position in the buffer.
func (b *Buffer) At(index int) rune {
	return b.buffer[index]
}

/*
	We need to present this buffer on the screen. Remember, a buffer can be really large, having
	hundrers of lines.
	We can only show a part of that buffer. For this we construct a viewport struct.
*/
type Viewport struct {
	buffer *Buffer
	offset int
}

func NewViewport(buffer *Buffer) *Viewport {
	return &Viewport{buffer: buffer}
}

func (v *Viewport) GetCharacterAt(index int) rune {
	return v.buffer.At(v.offset + index)
}

/*
	So, we would have lots of viewports and lots of buffers,
	but we also want a simple API for just creating a console which contains all of these constructs
	behind the scenes.
*/

type Console struct {
	buffers   []*Buffer
	viewports []*Viewport
	offset    int
}

// Let's make a factory function for initializing this console.
func NewConsole() *Console {
	// Initialized with a single buffer and a single viewport
	b := NewBuffer(200, 150)
	v := NewViewport(b)
	return &Console{[]*Buffer{b}, []*Viewport{v}, 0}
}

func (c *Console) GetCharacterAt(index int) rune {
	// By default, we will look into the first viewport
	return c.viewports[0].GetCharacterAt(index)
}

func main() {
	c := NewConsole()
	u := c.GetCharacterAt(1)
	fmt.Println(u)
}

/*
	Just because we are making a facade doesn't mean that we have to obscure the details.
	If somebody wanted to mess with the buffers and viewports, they can do it through the console.

*/
