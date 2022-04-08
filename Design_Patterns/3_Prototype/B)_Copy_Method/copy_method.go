package copymethod

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

/*
We can organize our own object to have a sort of deep copy method available.
This method can be invoked on any of the instances to perform a deep copying.

However, it still leaves open the problem of what to do with situations in which you have
a struct with slices, and all sorts of data structures.
In the case of the Friends slice, in this example, we cannot go ahead and add additional behaviours
to that slice.

Moreover, we would need to check every single one of the structs to make sure that everyone of the member
types has a deep copy method.

*/
func (p *Person) DeepCopy() *Person {
	q := *p
	q.Address = p.Address.DeepCopy()
	copy(p.Friends, q.Friends)
	return &q
}

func (a *Address) DeepCopy() *Address {
	return &Address{
		a.StreetAddress,
		a.City,
		a.Country,
	}
}
func main() {
	john := Person{
		"John",
		&Address{"123 London Rd.", "London", "UK"},
		[]string{"Chris", "Matt"},
	}
	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker St."
	jane.Friends = append(jane.Friends, "Angela")

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
