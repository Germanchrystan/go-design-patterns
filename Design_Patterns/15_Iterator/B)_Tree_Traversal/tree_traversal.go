package main

import "fmt"

/*
	The concept of an iterator in Go is not as common as it is in other programming languages,
	but we can take a look in an scenario in which it is in fact relevant, and where we cannot get
	away with without using an iterator.

	We are going to take a look at binary trees.
	A binary tree is a data structure which has a root with a value. This is the first node.
	Each node in a binary tree can have one, two or zero branches pointing to another node.

*/

type Node struct {
	Value               int
	left, right, parent *Node
}

func NewNode(value int, left, right *Node) *Node {
	n := &Node{Value: value, left: left, right: right}
	left.parent = n
	right.parent = n
	return n
}

func NewTerminalNode(value int) *Node {
	return &Node{Value: value}
}

/*
	root := NewNode(1,
		NewTerminalNode(2),
		NewTerminalNode(3))

					1
		         /	   \
				2		3
			What we want is to be able to traverse this binary tree.
			There are different algorithms, that apply different ways of traversing.
			The three most common ones are in order, preorder and post order.
			in-order:  213
			preorder:  123
			postorder: 231

			We are going to build an in-order iterator.
*/

type InOrderIterator struct {
	Current       *Node
	root          *Node
	returnedStart bool
	// This boolean variable indicates whether or not we return the starting value
}

func NewInOrderIterator(root *Node) *InOrderIterator {
	i := &InOrderIterator{
		root, root, false,
	}
	// We want to traverse the entire tree until we find the leftmost node
	for i.Current.left != nil {
		i.Current = i.Current.left
	}
	return i
}

func (i *InOrderIterator) Reset() {
	i.Current = i.root
	i.returnedStart = false
}

func (i *InOrderIterator) MoveNext() bool {
	if i.Current == nil {
		return false
	}
	if !i.returnedStart {
		i.returnedStart = true
		return true
	}
	if i.Current.right != nil {
		i.Current = i.Current.right
		for i.Current.left != nil {
			i.Current = i.Current.left
		}
		return true
	} else {
		p := i.Current.parent
		for p != nil && i.Current == p.right {
			i.Current = p
			p = p.parent
		}
		i.Current = p
		return i.Current != nil
	}
}

func main1() {
	root := NewNode(1,
		NewTerminalNode(2),
		NewTerminalNode(3))
	it := NewInOrderIterator(root)
	for it.MoveNext() {
		fmt.Printf("%d,", it.Current.Value)
	}
	fmt.Println("\b") //This will erase the comma in the last iteration
}

/*
	For a better implementation we could have a binary tree struct.
*/
type BinaryTree struct {
	root *Node
}

func NewBinaryTree(root *Node) *BinaryTree {
	return &BinaryTree{root: root}
}

func (b *BinaryTree) InOrder() *InOrderIterator {
	return NewInOrderIterator(b.root)
}

func main2() {
	root := NewNode(1,
		NewTerminalNode(2),
		NewTerminalNode(3))

	t := NewBinaryTree(root)
	for i := t.InOrder(); i.MoveNext(); {
		fmt.Printf("%d,", i.Current.Value)
	}
	fmt.Println("\b")
}

func main() {
	main1()
	main2()
}
