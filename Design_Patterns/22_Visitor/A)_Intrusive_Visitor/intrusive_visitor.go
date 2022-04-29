package main

import (
	"fmt"
	"strings"
)

/*
 	The first implementation, typically called intrusive representation.
	Whenever we call something intrusive, it means that it intrudes into
	the structure that is already created. Any intrusive approach is by
	definition a violation of the Open Closed Principle.

	We are going to stablish an 'Expression' interface to handle numeric expressions.
*/

type Expression interface {
	Print(sb *strings.Builder) // Being intrusive by adding a Print function
}

type DoubleExpression struct { // Using double precision floating point numbers
	value float64
}

func (d DoubleExpression) Print(sb *strings.Builder) {
	sb.WriteString(fmt.Sprintf("%g", d.value)) // %g formatting
}

type AdditionExpression struct {
	left, right Expression
}

func (a *AdditionExpression) Print(sb *strings.Builder) {
	sb.WriteRune('(')
	a.left.Print(sb)
	sb.WriteRune('+')
	a.right.Print(sb)
	sb.WriteRune(')')
}

func main() {
	// 1 + (2+3)
	e := AdditionExpression{
		left: &DoubleExpression{1},
		right: &AdditionExpression{
			left:  &DoubleExpression{2},
			right: &DoubleExpression{3},
		},
	}

	sb := strings.Builder{}
	e.Print(&sb)
	fmt.Println(sb.String())
	/*
		The visitor in this example is the string builder.
		It is the one that gets passed into every single print method.
		It gets to visit different expressions as well as double expressions,
		and that is why it's called visitor.
		It is not the best visitor because implementing this visitor implis modifying
		the behavior of both any interface that we have as part of the element hierarchy,
		as well as the elements themselves. So every single element suddenly has to have this
		additional method.
		Imagine that we want to have another visitor that actually calculates the value.
		Unfortunately in this setup, we should go on to the Expression interface,
		add another method to called evaluate, which returns float64,
		then implement this in both double expressions and additional expression.

		Another important concept that we need to cover here is the ideo of separation of
		concerns and single responsibility. It is kind of responsibility of each expression
		to print itself, but not necessarily. It would make more sense if we had a separate component,
		which knew how to print double expressions, additional expressions, as well as any other kind of
		expression added later to the program.
	*/
}
