package main

import (
	"fmt"
)

// To iterate something means to go through every or most of the elements that conform its structure.
type Person struct {
	FirstName, MiddleName, LastName string
}

// How would we go through every single name in the Person struct?
// One of the approaches is simply to expose an array

func (p *Person) Names() [3]string {
	return [3]string{p.FirstName, p.MiddleName, p.LastName}
}

func main1() {
	p := Person{"Alexander", "Graham", "Bell"}
	for _, name := range p.Names() {
		fmt.Println(name)
	}
}

/*
	Suppose that if the middle name is empty,
	we don't want to return that name as part of the Names function.

	Another approach is to use a generator, using a channel and a go routine
*/
func (p *Person) NamesGenerator() <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		out <- p.FirstName // This is all going to be running in parallel
		if len(p.MiddleName) > 0 {
			out <- p.MiddleName
		}
		out <- p.LastName
	}()
	return out

}

func main2() {
	p := Person{"Alexander", "Graham", "Bell"}
	for name := range p.NamesGenerator() {
		fmt.Println(name)
	}
}

/*
	The third is the most complicated variety of iteration.
	It is when we use a separate struct.
	This approach is very idiomatic. It is the kind of approach they use in C++.
*/

type PersonNameIterator struct {
	person  *Person
	current int
}

func NewPersonNameIterator(person *Person) *PersonNameIterator {
	return &PersonNameIterator{person, -1}
	// the current value is going to be -1, so when we start iterating, we move it to 0.
}

func (p *PersonNameIterator) MoveNext() bool {
	p.current++
	return p.current < 3 // Limiting the increase of current
}

func (p *PersonNameIterator) Value() string {
	switch p.current {
	case 0:
		return p.person.FirstName
	case 1:
		return p.person.MiddleName
	case 2:
		return p.person.LastName
	default:
		panic("We should not be here!")
	}
}

func main3() {
	p := Person{"Alexander", "Graham", "Bell"}
	for it := NewPersonNameIterator(&p); it.MoveNext(); {
		fmt.Println(it.Value())
	}
}
