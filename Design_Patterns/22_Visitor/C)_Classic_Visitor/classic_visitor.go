package main

import "strings"

/*
	Go doesn't support function overload.
	So this is not possible:

	func Print(e DoubleExpression, sb *strings.Builder) {
		sb.WriteString(fmt.Sprintf("%g", de.value))

	}

	func Print(e AdditionExpression, sb *strings.Builder) {
		sb.WriteRune('(')
		Print(e.left, sb)
		sb.WriteRune('+')
		Print(e.right, sb)
		sb.WriteRune(')')
	}

	============> Error: Print redeclared in this block

	Other reason why this would not work is because whenever the compiler encounters
	e.left or e.right, it knows that left or right are expressions. It doesn't know if
	they are additional expressions or double expressions, and it cannot figure it out at compilation time.
	The compiler wants to know the static type of e.left in order to be able to do something
	with it.

	This is the reason why the idea of double dispatch is used.
	Double dispatch is being able to choose the right method, not just on the basis
	of an argument, but also on the basis of who the caller is.

	The first thing to do is modify the Expression interface. The thing about the
	double dispatch classic visitor is that we can modify this interface, but only once.
	That operation is leveraged for many different kinds of visitors.
*/

type Expression interface {
	Accept(ev ExpressionVisitor)
}

type ExpressionVisitor interface {
	VisitDoubleExpression(e *DoubleExpression)
	VisitAdditionExpression(e *AdditionExpression)
}

//=======================================================================//
type DoubleExpression struct {
	value float64
}

func (d *DoubleExpression) Accept(ev ExpressionVisitor) {
	ev.VisitDoubleExpression(d)
}

//=======================================================================//
type AdditionExpression struct {
	left, right Expression
}

func (a *AdditionExpression) Accept(ev ExpressionVisitor) {
	ev.VisitAdditionExpression(a)
}

//=======================================================================//

type ExpressionPrinter struct {
	sb strings.Builder
}

func (ex ExpressionPrinter) VisitDoubleExpression(e *DoubleExpression) {

}

func (ex ExpressionPrinter) VisitAdditionExpression(e *AdditionExpression) {

}
