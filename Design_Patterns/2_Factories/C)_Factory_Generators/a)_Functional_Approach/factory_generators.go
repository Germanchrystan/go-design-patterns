package factorygenerators

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

/*
What we want to do is to be able to create factories dependent
upon the settings in which the Employee struct is manufactured
*/

/*
Functional approach
The functional approach would consist on creating a factory function that would not return an object,
but a function.
*/
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	// Very similar to clousures in JS
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

func main() {
	developerFactory := NewEmployeeFactory("Developer", 60000)
	managerFactory := NewEmployeeFactory("Manager", 80000)
	developer := developerFactory("Adam")
	manager := managerFactory("Jane")

	fmt.Println(developer)
	fmt.Println(manager)
}
