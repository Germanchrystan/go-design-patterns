package main

import (
	"fmt"
	"strings"
	"unicode"
)

type FormattedText struct {
	plainText  string
	capitalize []bool
}

func NewFormattedText(plainText string) *FormattedText {
	return &FormattedText{
		plainText:  plainText,
		capitalize: make([]bool, len(plainText)),
	}
}

func (f *FormattedText) String() string {
	sb := strings.Builder{}
	for i := 0; i < len(f.plainText); i++ {
		c := f.plainText[i]
		if f.capitalize[i] {
			sb.WriteRune(unicode.ToUpper(rune(c)))
		} else {
			sb.WriteRune(rune(c))
		}
	}
	return sb.String()
}

func (f *FormattedText) Capitalize(start, end int) {
	for i := start; i <= end; i++ {
		f.capitalize[i] = true
	}
}

func main() {
	text := "This is a brave new world"

	ft := NewFormattedText(text)
	ft.Capitalize(10, 15)
	fmt.Println(ft.String())
}

/*
	Unfortunately, this approach is extremely inefficient, because we are specifying a huge boolean slice,
	with one boolean for every single character inside the plain text.

	Imagine having the text of a long novel, and only wanting to capitalize a single word out of thousands
	of words. Lots of values are going to be allocated.

	We can introduce an idea of a text range.
*/

type TextRange struct {
	Start, End               int
	Capitalize, Bold, Italic bool
}

func (t *TextRange) Covers(position int) bool {
	return position >= t.Start && position <= t.End
}

type BetterFormattedText struct {
	plainText  string
	formatting []*TextRange
}

/*
	The reason why formatting is a slice of pointers is because we also want to share these text ranges.
	We want to be able to return them to the user to operate upon.
*/

func NewBetterFormattedText(plainText string) *BetterFormattedText {
	return &BetterFormattedText{plainText: plainText}
}

func (b *BetterFormattedText) Range(start, end int) *TextRange {
	r := &TextRange{start, end, false, false, false}
	b.formatting = append(b.formatting, r)
	return r // by returning the range pointer, the user can operate upon it.
}

func (b *BetterFormattedText) String() string {
	sb := strings.Builder{}
	for i := 0; i < len(b.plainText); i++ {
		c := b.plainText[i]
		for _, r := range b.formatting {
			if r.Covers(i) && r.Capitalize {
				c = uint8(unicode.ToUpper(rune(c)))
			}
		}
		sb.WriteRune(rune(c))
	}
	return sb.String()
}

func main2() {
	text := "This is a brave new world"

	bft := NewBetterFormattedText(text)
	bft.Range(16, 19).Capitalize = true
	fmt.Println(bft.String())
}
