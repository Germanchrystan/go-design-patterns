package interfacefactory

import "fmt"

type Person interface {
	SayHello()
}

/*
When you have a factory function, you don't always have to return a struct.
Instead, what can be done is to create a factory that returns an interface that the struct conforms to
*/
type person struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Printf("Hi, my name is %s, I am %d years old \n", p.name, p.age)
}

/*
	This is a good way to encapsulate information.
	In this case, the underlying type struct can not be accesible, since it is in lowercase,
	as well as the struct properties.
	The struct is only accesible through its interface
	p := NewPerson("James", 35)
	p.age++ =>  would not work.
*/
func NewPerson(name string, age int) Person /* Returning an interface */ {
	return &person{name, age}
}

// Other structs could conform to this interface
type tiredPerson struct {
	name string
	age  int
}

func (p *tiredPerson) SayHello() {
	fmt.Printf("Sorry, I am too tired to talk to you\n")
}

/*
This is a different way of implementing NewPerson, in which depending on the age argument,
one of two structs is returned. Both conform to the Person interface.
*/
func AnotherImplementationOfNewPerson(name string, age int) Person {
	if age > 100 {
		return &tiredPerson{name, age}
	}
	return &person{name, age}
}

func main() {
	p := NewPerson("James", 35)
	p.SayHello()

}
