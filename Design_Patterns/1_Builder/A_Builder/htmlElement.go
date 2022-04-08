package firstexample

import (
	"fmt"
	"strings"
)

const (
	indentSize = 2
)

// The entire HTML construct is a tree.
type HtmlElement struct {
	name, text string
	elements   []HtmlElement // Nested Elements
}

func (e *HtmlElement) String() string {
	return e.string(0)
}

func (e *HtmlElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.name))
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ", indentSize*(indent+1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}

	for _, el := range e.elements {
		sb.WriteString(el.string(indent + 1)) // Recursion
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.name))
	return sb.String()
}

type HtmlBuilder struct {
	rootName string
	root     HtmlElement
}

func NewHtmlBuilder(rootName string) *HtmlBuilder {
	return &HtmlBuilder{rootName, HtmlElement{rootName, "", []HtmlElement{}}}
}

func (b *HtmlBuilder) String() string {
	return b.root.String()
}

func (b *HtmlBuilder) AddChild(childName, childText string) {
	e := HtmlElement{childName, childText, []HtmlElement{}}
	b.root.elements = append(b.root.elements, e)
}

// Adding the return it is possible to concatenate calls to this function
func (b *HtmlBuilder) AddChildFluent(childName, childText string) *HtmlBuilder {
	e := HtmlElement{childName, childText, []HtmlElement{}}
	b.root.elements = append(b.root.elements, e)
	return b
}

func main_() {
	b := NewHtmlBuilder("ul")
	// Adding elements
	b.AddChild("li", "hola")
	b.AddChild("li", "mundo")
	// Chaining calls with fluent interface
	b.AddChildFluent("li", "hello").AddChildFluent("li", "world")
}
