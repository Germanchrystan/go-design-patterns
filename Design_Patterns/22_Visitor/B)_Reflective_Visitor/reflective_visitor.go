package main

import (
	"fmt"
	"strings"
)

/*
	Let's take the previous example again.
	Imagine we wanted to concentrate all the different methods in a
	separate component, whether it is a struct of a function.

	We are going to build a very simple reflective visitor.
	Why is it reflective?
	Typically there is this construct called reflexion.
	That is when we look into a type and what the type actually is,
	what kind of members it has, and Go has certain amount of support for reflexion.

	One of the trademarks of reflection is that it checks the actual type.
	In the print method we have a Expression type argument, but we do not know what kind of
	Expression it is.
*/
func Print(e Expression, sb *strings.Builder) {
	// Checking the type
	if de, ok := e.(*DoubleExpression); ok { //If the casting is successful, ok will equal true
		sb.WriteString(fmt.Sprintf("%g", de.value))
	} else if ae, ok := e.(*AdditionExpression); ok {
		sb.WriteRune('(')
		Print(ae.left, sb)
		sb.WriteRune('+')
		Print(ae.right, sb)
		sb.WriteRune(')')
	}
}

// Now we have a self-contained function specifically for printing the expression.

type Expression interface {
}

type DoubleExpression struct {
	value float64
}

type AdditionExpression struct {
	left, right Expression
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
	Print(e, &sb)
	fmt.Println(sb.String())
	/*
		This approach is better than the previous one.
		It is better because we have taken out this particular concern, the printing concern.
		The most obvious downside is what will happen if there is a third type. Let's say we
		introduce substraction as a new type of expression. All of the sudden we have to write
		additional code in the Print function. We are still breaking the Open Closed Principle.
		Imagine if we forgot to add this case to the Print function, it would not be even considered.
	*/
}
