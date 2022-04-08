package copythroughserialization

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// Adding support for encoding

/*

We are going to be using binary serialization.
Serialization constructs are very smart. A serializaer is going
to figure out each type of the members of a struct that needs to be deep copied.
Even the values that are accesible through pointers.

A serializer knows how to unwrap a structure and serialize all of its members.
If the struct person is serialized to a file, its state is saved, including the dependencies.
When the struct is de-serialized, we construct a brand new object initialized with the same values.
*/

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	b := bytes.Buffer{}     // Creating a buffer
	e := gob.NewEncoder(&b) // e is going to be an encoder, which takes a pointer to the buffer
	_ = e.Encode(p)         // Encoding the person instance

	fmt.Println(string(b.Bytes())) // This is just written to take a look into the encoded object
	// Some of the characters will be buggy, since some of it won't be printable

	d := gob.NewDecoder(&b)
	result := Person{}
	_ = d.Decode(&result) // Decoding into a new instance of Person
	return &result
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
