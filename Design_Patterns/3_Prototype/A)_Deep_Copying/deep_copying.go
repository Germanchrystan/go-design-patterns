package deepcopying

/*
	Deep Copying
	It consist on making copies of an object with everything it refers to,
	including pointers, slices, etc.
	When you don't make a copy of pointers between objects you are having the same pointer
	as a property of two or more objects, and that prevents customization.
*/

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
}

func main() {
	john := Person{"John", &Address{"123 London Rd", "London", "UK"}}
	/*
		Let's say we want to create another instance of person that lives at the same address,
		but have different names.
	*/
	jane := john
	jane.Name = "Jane" // This is ok.
	/*
		The problem now is that both of this instances have pointers of the same address instance.
		Address won't be able to be customized later, because is the same address for John

		jane.Address.StreetAddress = "321 Baker St"

		How do we fix this? Well, jane's address should be a different pointer
		jane.Address = &Address{
			john.Address.StreetAddress,
			john.Address.City,
			john.Address.Country
		}

		Now, Jane's address can be customized.
		However, this approach doesn't scale.
	*/
	//---------------------------------------------------------------------------------------------//
}
