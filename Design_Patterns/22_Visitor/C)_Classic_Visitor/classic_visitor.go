package main

import (
	"fmt"
	"strings"
)

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

func NewExpressionPrinter() *ExpressionPrinter {
	return &ExpressionPrinter{strings.Builder{}}
}

func (ep *ExpressionPrinter) VisitDoubleExpression(e *DoubleExpression) {
	ep.sb.WriteString(fmt.Sprintf("%g", e.value))
}

func (ep *ExpressionPrinter) VisitAdditionExpression(e *AdditionExpression) {
	ep.sb.WriteRune('(')
	e.left.Accept(ep)
	ep.sb.WriteRune('+')
	e.right.Accept(ep)
	ep.sb.WriteRune(')')

	/*
		In the Accept method is where the double dispatch magic happens.
		So we call e.left.Accept(), and the reason why we call this,
		regardless of what left is, is because e.left is an expression,
		and an expression is an interface that defines a method called Accept.
		So we know that the method Accept is there, and we pass the ep as the argument.
		So, we end up going to the Accept method, which in turn returns us to one
		of the Visit Expression functions. By doing this double jump, we are able
		to have all the information about the callee and the caller.
	*/
}

func (ep *ExpressionPrinter) String() string {
	return ep.sb.String()
}

func main1() {
	// 1 + (2+3)
	e := AdditionExpression{
		left: &DoubleExpression{1},
		right: &AdditionExpression{
			left:  &DoubleExpression{2},
			right: &DoubleExpression{3},
		},
	}
	ep := NewExpressionPrinter()
	e.Accept(ep)
	fmt.Println(ep.String())

	/*
		Let's talk about the extensibility of this approach,
		because we have been fighting for having the support of the Open Closed Principle.
		There are still some modifications that might be required.
		Imagine adding a Substraction expression.
		We would have to add a VisitSubstractionExpression to the ExpressionVisitor interface.
		But as soon as we do this all of the visitors involved would have to support substraction
		expressions. This is a drastic difference, because in this case we can't forget to handle this case.
	*/
}

// Imagine we want to add an Evaluate expression
type ExpressionEvaluator struct {
	result float64
}

func (ee ExpressionEvaluator) VisitDoubleExpression(e *DoubleExpression) {
	ee.result = e.value
}

func (ee ExpressionEvaluator) VisitAdditionExpression(e *AdditionExpression) {
	e.left.Accept(ee)
	x := ee.result
	e.right.Accept(ee)
	x += ee.result
	ee.result = x
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
	ep := NewExpressionPrinter()
	e.Accept(ep)

	ee := &ExpressionEvaluator{}
	e.Accept(ee)
	fmt.Printf("%s = %g", ep, ee.result)
}
