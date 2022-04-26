# Iterator Design Pattern
How traversal of data structures happens and who makes it happen.

The iterator design pattern that facilitates the traversal of data structures.
- Iteration (traversal) is a core functionality of various data structrures.
- An iterator is a type that facilitates the traversal.
--- Keeps a pointer to the current element
--- Knows how to move to a different element. 

- Go allows iteration with range
--- Built in support in many objects(arrays, slices, etc.)
--- Sometimes we need to support iterations in our own structures as well.

### In Summary
- An iterator specifies how to traverse an object.
- Moves along the iterated collection, indicating when the last element has been reached.
- Not particularly idiomatic in Go: There is no standard iterable interface, but they are still easy to build and use.