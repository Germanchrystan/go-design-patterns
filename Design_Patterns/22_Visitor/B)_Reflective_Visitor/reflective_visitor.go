package main

import "strings"

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
*/
func Print(e Expression, sb *strings.Builder) {}

type Expression interface {
}

type DoubleExpression struct {
	value float64
}

type AdditionExpression struct {
	left, right Expression
}
