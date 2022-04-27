package main

import (
	"container/list"
	"fmt"
)

/*
	One very common use of the observer design pattern is to implement notifications
	related to property changes on an object.

	Let's suppose we want to be informed when a person's age changes.
*/

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

/*
	We could define a Getter and Setter for a Person's age.
	These concepts are not very idiomatic to Go, unless we want to do
	the kind of thing we are doing here, which is change notifications.
	When we set the age, we will also notify the subscribers.

	We can have a separate struct for encoding the information necessary in a notification

*/
type PropertyChange struct {
	Name  string      // "Age" or any other property
	Value interface{} // Generic
}

func (p *Person) Age() int { return p.age }

func (p *Person) SetAge(age int) {
	if age == p.age {
		return
	}
	p.age = age
	p.Fire(PropertyChange{"Age", p.age})
}

//===========================================================//
/*
	A traffic management company wantd to be informed about a person's age.
	If the person is too young to drive, the company would keep monitoring.
	As soon as the person turns 16, the traffic management congratulates the person
	on being able to get a license and then it unsubscribes from the observable.
*/
type TrafficManagement struct {
	o Observable
}

func (t *TrafficManagement) Notify(data interface{}) {
	// Checking is the data argument is of type PropertyChange, just to be safe.
	if pc, ok := data.(PropertyChange); ok {
		if pc.Value.(int) >= 16 {
			fmt.Println("Congrats, you can drive now!")
			t.o.Unsubscribe(t)
		}
	}
}

//===========================================================//

func main() {
	p := NewPerson(15)
	// We pass the observable part of the person
	t := &TrafficManagement{p.Observable}
	p.Subscribe(t)

	for i := 16; i <= 20; i++ {
		fmt.Println("Setting the age to", i)
		p.SetAge(i)
	}
}

/*
	This approach does have certain problems with dependencies,
	and that is what we are going to take a look at in the next example.
*/
