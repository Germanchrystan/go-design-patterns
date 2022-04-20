package main

import (
	"fmt"
	"strings"
)

/*
	We are going to take a look at what is probablyu the most classic implementation
	of the flyweight desing pattern.

*/

type User struct {
	FullName string
}

func NewUser(fullname string) *User {
	return &User{FullName: fullname}
}

func main1() {
	john := NewUser("John Doe")
	jane := NewUser("Jane Doe")
	alsoJane := NewUser("Jane Smith")

	fmt.Println("Memory taken by users: ",
		len([]byte(john.FullName))+len([]byte(alsoJane.FullName))+len([]byte(jane.FullName)))
	/*
			In the real world, if you take thousands and thousands of different players in a system,
		 	you are going to have lots of people with the same names. And every time you store one
			of those repeated names, you are duplicating memory.

			Why can't we somehow compact this data, store it inside a single container, and then have
			indices into the names contained within that structure?
	*/
}

var allNames []string

type User2 struct {
	names []uint8 // All of the parts of a full name will become integers
}

func NewUser2(fullname string) *User2 {

	getOrAdd := func(s string) uint8 {
		for i := range allNames {
			if allNames[i] == s {
				// Then we return the index casted to a uint8
				return uint8(i) // Assuming that there is only going to be 256 unique names at most.
			}
		}
		allNames = append(allNames, s)
		return uint8(len(allNames) - 1)
	}
	result := User2{}
	parts := strings.Split(fullname, " ")

	for _, p := range parts {
		result.names = append(result.names, getOrAdd(p))
	}
	return &result
}

func (u *User2) FullName() string {
	var parts []string
	for _, id := range u.names {
		parts = append(parts, allNames[id])
	}
	return strings.Join(parts, " ")
}

func main2() {
	john2 := NewUser2("John Doe")
	jane2 := NewUser2("Jane Doe")
	alsoJane2 := NewUser2("Jane Smith")
	totalMem := 0

	/*
		Here, the memory calculation is different because we have to take
		all the names inside the system first and to calculate their total length.

		Then, we also add the length of every single one of those byte arrays.

	*/
	for _, a := range allNames {
		totalMem += len([]byte(a))
	}
	totalMem += len(john2.names) /* If these were integer rate, we would have to multiply this by some value
	like four or byte, but here we are using uint8, so there is no need for further calculation. */
	totalMem += len(jane2.names)
	totalMem += len(alsoJane2.names)
	fmt.Println("Memory taken by users2: ", totalMem) // Less memory than the previous model
}

func main() {
	main1()
	main2()
}
