package factorygenerators

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

/*
	Structured Approach
It consist in making a factory a struct.
*/
type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.AnnualIncome}
}

// Then, what would be required is to have predefined employee factories.
func NewEmployeeFactory(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}

/*
The only real advantages that the structured approach has over the functional approach
is that the functional factories can not be later modified. You can not change the income of the developer,
for instance. In the structured approach, default properties can be modified, since they are fields
	bossFactory.AnnualIncome = 110000

On the other hand, in terms of usability in third party code, providing a function and passing it into
some other piece of API is easier than passing a specialized object, for which the developer has to know
that it has a "Create" method, for instance. For this last situation, you might try to introduce some sort of
interface, which tells explicitly that there is a "Create" method to be used.

Both options are valid eitherway.
*/

func main() {
	bossFactory := NewEmployeeFactory("CEO", 100000)
	boss := bossFactory.Create("Sam")
	fmt.Println(boss)
}
