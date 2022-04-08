package firstexample

import (
	"fmt"
	"strings"
)

/*
First of, we are going to begin by using a builder that is actually already built into Go, and that is the string builder
A web server is supposed to serve HTML. It also serves other things like JavaScript.
The idea is that you need to build up string of HTML from ordinary text elements.

*/

func main() {
	// So, for example, you have a piece of text and you want to turn it into a paragraph.
	hello := "Hello"
	//strings.Builder is a built in component, which is part of the Go SDK, and actually helps to concatenate strings together.
	stringBuilder := strings.Builder{}
	stringBuilder.WriteString("<p>")
	stringBuilder.WriteString(hello)
	stringBuilder.WriteString("</p>")

	fmt.Println(stringBuilder.String()) // => <p>Hello</p>

	// Let's suppose you have a list of words and you want to put them into a HTML unordered list.
	words := []string{"hello", "world"}
	stringBuilder.Reset()
	stringBuilder.WriteString("<ul>")
	for _, v := range words {
		stringBuilder.WriteString("<li>")
		stringBuilder.WriteString(v)
		stringBuilder.WriteString("</li>")
	}
	stringBuilder.WriteString("</ul>")

	// This process above is a bit too complicated. So it would be better to put it into structures.
	// In htmlElement.go we will see how to do this.
}
