package firstexample

import "fmt"

// Enterprise pattern: Specification

/*
Let's imagine that you are operating some sort of online store.
You are selling widgets of some kind, just real physical objects.
And let's suppose that you want the end usr of your website to be able to filter those items by size or color.
*/
type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

// Filter
type Filter struct {
	//
}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

// Imagine that now you need to filter by size
// You now have to add a method to filter by size
func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

// Now filter by size AND color is required
// You have to now add a third method
func (f *Filter) FilterBySizeAndColor(products []Product, size Size, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size && v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

// This is a violation of the OCP.
// Filter struct is not open for extension, cause we need to modify it every time a new filter is required.

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}

	// Filtering by color
	fmt.Printf("Green products: \n")
	f := Filter{}
	for _, v := range f.FilterByColor(products, green) {
		fmt.Printf(" - %s is green \n", v.name)
	}
}
