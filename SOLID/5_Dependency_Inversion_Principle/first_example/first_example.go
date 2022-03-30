package firstexample

import "fmt"

/*
Let's suppose you are doing some sort of genealogy research and you want to model
relationships between different people.
*/

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

// (Low level module: it has a storage functionality)
type Relationships struct {
	relations []Info
}

// We need to be able to add those kinds of relationships
func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(
		r.relations,
		Info{parent, Parent, child})
	r.relations = append(
		r.relations,
		Info{child, Child, parent})
}

// We also want to perform research
// (High level module: it operates on data)
type Research struct {
	// Here we have a high level module depending on a low level module
	relationships Relationships
}

func (r *Research) Investigate() {
	// In order to be able to actually perform any research, we have to go into relantionships,
	// and look at them.
	relations := r.relationships.relations // Getting convoluted
	for _, rel := range relations {
		if rel.from.name == "John" && rel.relationship == Parent {
			fmt.Println("John has a child called ", rel.to.name)
		}
	}
}

func main() {
	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Matt"}

	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	r := Research{relationships}
	r.Investigate()
}

/*
Everything looks ok. But there is a major problem with this entire scenario,
because what is happening is the research module here is actually using the internals of the
relationships module. So relationships is a low level module and Research is using its slice to get data from it.

If relationships decides to change the storage mechanic from a slice to a database, Research must be reworked.
In fact, it could be argued that the finding of a child od a particular person is something that needs to be handled not in a
high level module, but in a low level module. Essentially, if you know the storage mechanic, you can do an optimized research.
Queries can be made at a low level.
*/
