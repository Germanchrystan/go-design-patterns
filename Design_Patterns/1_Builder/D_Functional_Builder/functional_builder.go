package functionalbuilder

import "fmt"

/*
One way of extending a builder is by using a functional programming approach.
*/

type Person struct {
	name, position string
}

type personMod func(*Person)
type PersonBuilder struct {
	actions []personMod
}

func (b *PersonBuilder) Called(name string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.name = name
	})
	return b
}

func (b *PersonBuilder) Build() *Person {
	p := Person{}
	for _, a := range b.actions {
		a(&p)
	}
	return &p
}

func (b *PersonBuilder) WorkAsA(position string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.position = position
	})
	return b
}

func main() {
	b := PersonBuilder{}
	p := b.Called("Dimitri").WorkAsA("Developer").Build()
	fmt.Println(*p)
}
