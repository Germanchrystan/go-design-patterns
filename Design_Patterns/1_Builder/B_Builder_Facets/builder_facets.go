package builderfacets

import "fmt"

/*
In most situations that you encounter in daily programming, a single builder is sufficient for building a particular object.
But there are situations where you need more  than one builder, and it is necessary to somehow separate the process of building up the
different aspects of a particular type
*/

type Person struct {
	// address
	StreetAddress, Postcode, City string
	// job
	CompanyName, Position string
	AnnualIncome          int
}

type PersonBuilder struct {
	person *Person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

// We can have additional builder for the address and the job information
type PersonAddressBuilder struct {
	PersonBuilder
}

type PersonJobBuilder struct {
	PersonBuilder
}

// For person builder, we want to be able to provide interfaces which are given by the address builder and the job builder respectively.
func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{
		*b, // Here we make a copy of the builder, we put it into an address builder and that copies the pointer.
	}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

/*
So now we have ways of transitioning from a person builder to either a person address builder or person job builder.
But we need to realize that effectively, person builder and person address builder are both person builbers.
And as a result, when we have a person address builder, we can jump to a person job builder and viceversa.
*/

/*
We can populate the methods of the person job builder and personal address builder
*/
//---------------------------------------------------------------------------------//

func (b *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	b.person.StreetAddress = streetAddress
	return b
}

func (b *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	b.person.City = city
	return b
}

func (b *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	b.person.Postcode = postcode
	return b
}

//---------------------------------------------------------------------------------//
func (b *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	b.person.CompanyName = companyName
	return b
}

func (b *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	b.person.Position = position
	return b
}

func (b *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
	b.person.AnnualIncome = annualIncome
	return b
}

//---------------------------------------------------------------------------------//

func (b *PersonBuilder) Build() *Person {
	return b.person
}

//---------------------------------------------------------------------------------//

func main() {
	pb := NewPersonBuilder()
	pb.
		Lives().At("123 London Road").In("London").WithPostcode("SW12BC").
		Works().At("Fabrikam").AsA("Programmer").Earning(123000)
	person := pb.Build()
	fmt.Println(person)
}
