package main

import (
	"container/list"
	"fmt"
)

/*
	Let's imagine that we have a situation where there is a person that becomes ill.
	And when that happens, be want the doctor to be notified, so they can visit the patient.
*/

type Observable struct {
	subs *list.List //List of subscribers. Using the list library would work just fine.
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

func (o *Observable) Fire(data interface{}) { // Using interface{} because we don't know the data type
	for z := o.subs.Front(); z != nil; z = z.Next() {
		z.Value.(Observer).Notify(data)
	}
}

// Observer Interface
type Observer interface {
	Notify(data interface{})
}

// Observable: Patient
type Person struct {
	Observable // Composition
	Name       string
}

//=============================================================//
func NewPerson(name string) *Person {
	return &Person{
		Observable: Observable{new(list.List)},
		Name:       name,
	}
}

func (p *Person) CatchACold() {
	p.Fire(p.Name)
}

//=============================================================//
type DoctorService struct {
}

func (d *DoctorService) Notify(data interface{}) {
	fmt.Printf("A doctor has been called for %s", data.(string)) // Data casted to a string
}

//=============================================================//
func main() {
	p := NewPerson("Boris")
	ds := &DoctorService{}
	p.Subscribe(ds) //We get the doctor service to subscribe to events which happen on the person
	p.CatchACold()
}

//=============================================================//
