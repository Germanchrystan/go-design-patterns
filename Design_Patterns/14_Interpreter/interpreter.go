package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Element interface {
	Value() int
}

type Integer struct {
	value int
}

func (i *Integer) Value() int {
	return i.value
}

func NewInteger(value int) *Integer {
	return &Integer{value: value}
}

type Operation int

const (
	Addition Operation = iota
	Substraction
)

type BinaryOperation struct {
	Type        Operation
	Left, Right Element
}

func (b *BinaryOperation) Value() int {
	switch b.Type {
	case Addition:
		return b.Left.Value() + b.Right.Value()
	case Substraction:
		return b.Left.Value() - b.Right.Value()
	default:
		panic("Unsupported operation")
	}
}

/*
	We are going to assume that every expression is a binary operation here.

	We need to define a parse function, that is going to turn the set of tokens
	into a top level binary operation.

*/
func Parse(tokens []Token) Element {
	result := BinaryOperation{}
	haveLhs := false
	for i := 0; i < len(tokens); i++ {
		token := &tokens[i]
		switch token.Type {
		case Int:
			n, _ := strconv.Atoi(token.Text)
			integer := Integer{n}
			if !haveLhs {
				result.Left = &integer
				haveLhs = true
			} else {
				result.Right = &integer
			}
		case Plus:
			result.Type = Addition
		case Minus:
			result.Type = Substraction
			/*
				Now we are going to deal with the most complicated part of all:
				the left and right parentheses.

				We are going to do the following:
				When we encounter a left partentheses, we are going to find the location where the right
				parentheses is located, we are going to take everything in between, and then we are going
				to feed it recursively into the parse method.
			*/
		case Lparen:
			j := i
			for ; j < len(tokens); j++ {
				if tokens[j].Type == Rparen {
					break
				}
			}
			var subExp []Token
			for k := i + 1; k < j; k++ {
				subExp = append(subExp, tokens[k])
			}
			element := Parse(subExp)
			if !haveLhs {
				result.Left = element
				haveLhs = true
			} else {
				result.Right = element
			}
			i = j
		}
	}
	return &result
}

type TokenType int

const (
	Int TokenType = iota
	Plus
	Minus
	Lparen
	Rparen
)

type Token struct {
	Type TokenType
	Text string
}

func (t *Token) String() string {
	return fmt.Sprintf("`%s`", t.Text)
}

func Lex(input string) []Token {
	var result []Token

	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '+':
			result = append(result, Token{Plus, "+"})
		case '-':
			result = append(result, Token{Minus, "-"})
		case '(':
			result = append(result, Token{Lparen, "("})
		case ')':
			result = append(result, Token{Rparen, ")"})
		default:
			sb := strings.Builder{}
			for j := i; j < len(input); j++ {
				if unicode.IsDigit(rune(input[j])) {
					sb.WriteRune(rune(input[j]))
					i++
				} else {
					result = append(result, Token{Int, sb.String()})
					i--
					break
				}
			}
		}
	}
	return result
}

func main() {
	input := "(13+4)-(12+1)"
	tokens := Lex(input)
	parsed := Parse(tokens)
	fmt.Printf("%s = %d\n", input, parsed.Value())
}
