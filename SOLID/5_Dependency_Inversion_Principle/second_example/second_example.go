package secondexample

import "fmt"

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
	// ...
}

type Info struct {
	from         *Person      // John...
	relationship Relationship // ...is a parent...
	to           *Person      // ...of Jill.
}

// We create a new low level module called RelationshipBrowser
type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

type Relationships struct {
	relations []Info
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)

	for i, v := range r.relations {
		if v.relationship == Parent && v.from.name == name {
			result = append(result, r.relations[i].to)
		}
	}
	return result
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(
		r.relations,
		Info{parent, Parent, child})
	r.relations = append(
		r.relations,
		Info{child, Child, parent})
}

// Here we introduce a new definition of Research
type Research struct {
	//Now the high level module depends on an abstraction
	browser RelationshipBrowser
}

func (r *Research) Investigate() {
	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called ", p.name)
	}
}

func main() {
	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Matt"}

	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	/*
		Now, if you decide to change the store mechanic of relations from slices to something else,
		then you would only be modifying the methods of relationships.
		You wouldn't have to change the methods of research, because it does not depend on the
		low level details.
	*/
	r := Research{&relationships}
	r.Investigate()
}
