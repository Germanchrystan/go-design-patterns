package firstexample

import "fmt"

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
type BetterFilter struct {
	//
}

func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

// Enterprise pattern: Specification
/*
In this example, the enterprise pattern will also be introduced,
because it is a really good way to illustrate this principle.
*/
type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

type SizeSpecification struct {
	size Size
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == s.size
}

// Composite specification for Size and Color
type AndSpecification struct {
	first, second Specification
}

func (a AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) &&
		a.second.IsSatisfied(p)
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}

	// Filtering by color
	fmt.Printf("Green products: \n")
	greenSpec := ColorSpecification{green}
	bf := BetterFilter{}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green \n", v.name)
	}
	largeSpec := SizeSpecification{large}
	lgSpec := AndSpecification{greenSpec, largeSpec}
	fmt.Printf("Large green products: \n")
	for _, v := range bf.Filter(products, lgSpec) {
		fmt.Printf(" %s is large and green", v.name)
	}
}
