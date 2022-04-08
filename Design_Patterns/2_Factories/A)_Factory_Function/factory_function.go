package factoryfunction

type Person struct {
	Name string
	Age  int
	// In 99.99% of the cases, a person has 2 eyes. We would like to give this property a default value
	EyeCount int
}

// Factory function
// A factory function is nothing more than a freestanding function
// which returns an instance of the struct you want to create.
func NewPerson(name string, age int) *Person {
	// This function could also implement validation
	if age < 16 {
		// ...
	}
	return &Person{name, age, 2}
}

func main() {
	p := NewPerson("John", 33)
	// Object could be customized later
	p.EyeCount = 1
}
