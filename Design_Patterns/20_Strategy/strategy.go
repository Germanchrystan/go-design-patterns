package main

import (
	"fmt"
	"strings"
)

type OutputFormat int

const (
	Markdown OutputFormat = iota
	Html
)

/*
	The idea is that whenever we print a list, there are typically
	three things we need to know about how to print it.
	There is a start of a list (in Html <ul>).
	Then we have every single list item.
	Finally there is the closing of the list. (</ul>)
*/
type ListStrategy interface {
	Start(builder *strings.Builder)
	End(build *strings.Builder)
	AddListItem(build *strings.Builder, item string)
}

//===================================================================================//
type MarkdownListStrategy struct {
}

func (m MarkdownListStrategy) Start(builder *strings.Builder) {
	// In Markdown there is no initialization of lists, like in Html
}
func (m MarkdownListStrategy) End(builder *strings.Builder) {
	// In Markdown there is no finalization of lists, like in Html
}
func (m MarkdownListStrategy) AddListItem(builder *strings.Builder, item string) {
	builder.WriteString(" * " + item + "\n")
}

//===================================================================================//
type HtmlListStrategy struct {
}

func (m HtmlListStrategy) Start(builder *strings.Builder) {
	builder.WriteString("<ul>\n")
}
func (m HtmlListStrategy) End(builder *strings.Builder) {
	builder.WriteString("</ul>\n")
}
func (m HtmlListStrategy) AddListItem(builder *strings.Builder, item string) {
	builder.WriteString("	<li>" + item + "</li>\n")
}

//===================================================================================//
type TextProcessor struct {
	builder      strings.Builder
	listStrategy ListStrategy
}

func NewTextProcessor(listStrategy ListStrategy) *TextProcessor {
	return &TextProcessor{
		builder:      strings.Builder{},
		listStrategy: listStrategy,
	}
}

func (t *TextProcessor) SetOutputFormat(fmt OutputFormat) {
	switch fmt {
	case Markdown:
		t.listStrategy = &MarkdownListStrategy{}

	case Html:
		t.listStrategy = &HtmlListStrategy{}
	}

}

func (t *TextProcessor) AppendList(items []string) {
	s := t.listStrategy
	s.Start(&t.builder)
	for _, item := range items {
		s.AddListItem(&t.builder, item)
	}
	s.End(&t.builder)
}

func (t *TextProcessor) Reset() {
	t.builder.Reset()
}

func (t *TextProcessor) String() string {
	return t.builder.String()
}

func main() {
	tp := NewTextProcessor(&MarkdownListStrategy{})
	tp.AppendList([]string{"foo", "bar", "baz"})
	fmt.Println(tp)

	tp.Reset()
	tp.SetOutputFormat(Html)
	tp.AppendList([]string{"foo", "bar", "baz"})
	fmt.Println()
}
