package main

import (
	"container/list"
	"fmt"
)

/*
	There are certain problems with using change notifications on properties.
	By property, we mean a combination of a getter and setter.

	Imagine that for a given person, we decide to make a read-only property.
	This property is computed from some other value, as opposed to just being used directly.
	For example, we want a boolean method telling us whether a person can vote.
*/

func (p *Person) CanVote() bool {
	return p.age >= 18
}

/*
	The problem is that we would want change notification for this as well.
	Change notifications can only happen in setters.
	It cannot really happen in Getters like CanVote.
	Where exactly would we send the notification for the change in the voting status?
	We would do this in the setter for the age.
*/

type ElectoralRoll struct{}

func (e *ElectoralRoll) Notify(data interface{}) {
	if pc, ok := data.(PropertyChange); ok {
		if pc.Name == "CanVote" && pc.Value.(bool) {
			fmt.Println("Congratulations, you can vote!")
		}
	}
}

/*
	Where would the CanVote event be fired? In the Setter, there is no other place for us to do it.
	In order to be able to do this, we need to keep track of changes in the CanVote boolean.
	So, we need to cache the previous value of the value and compare it to the new one,
	everytime the age is setted.
*/

func (p *Person) SetAge(age int) {
	if age == p.age {
		return
	}

	oldCanVote := p.CanVote() // Cache

	p.age = age
	p.Fire(PropertyChange{"Age", p.age})

	if oldCanVote != p.CanVote() {
		p.Fire(PropertyChange{"CanVote", p.CanVote()})
	}
}

func main() {
	p := NewPerson(0)
	er := &ElectoralRoll{}
	p.Subscribe(er)

	for i := 10; i < 20; i++ {
		fmt.Println("Setting age to", i)
		p.SetAge(i)
	}
}

/*
	The problem here is dependencies.
	The age setter becomes very large. It starts caching previous values of all the
	properties it affects, and then compares the previous values of
	the affected properties to the current ones.
	This situation doesn't really scale.
	This example here is just a very simple illustration of how we can get complexity
	virtually out of nowhere, and suddenly there is no Golang mechanism that would help us with it.
	If you wanted to have changed notifications on dependent properties, you would have to build
	some sort of complex infrastructure.
*/

//=======================================================================//
type Observable struct {
	subs *list.List
}

func (o *Observable) Subscribe(x Observer) {
	o.subs.PushBack(x)
}

func (o *Observable) Unsubscribe(x Observer) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		if z.Value.(Observer) == x {
			o.subs.Remove(z)
		}
	}
}

func (o *Observable) Fire(data interface{}) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		z.Value.(Observer).Notify(data)
	}
}

type Observer interface {
	Notify(data interface{})
}

//===========================================================//
type Person struct {
	Observable
	Name string
	age  int
}

func NewPerson(age int) *Person {
	return &Person{
		Observable: Observable{new(list.List)},
		age:        age,
	}
}

type PropertyChange struct {
	Name  string      // "Age" or any other property
	Value interface{} // Generic
}

func (p *Person) Age() int { return p.age }

//=======================================================================//
