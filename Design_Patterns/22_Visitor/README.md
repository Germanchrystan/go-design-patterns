# Visitor Design Pattern
*Allows adding extra behaviors to entire hierarchies of types*

* One of the problems that we sometimes try to solve is to define a new operation, not on a single type, but on several types, or on an entire type hierarchy at the same time.
    * E.g., given a document model (lists, paragraphs,etc.), we want to add printing functionality.
*Do not want to keep modifying every type in the hierarchy.
* Want to have the new functionality separate (Separate Responsibility Principle).
* This approach is often used for traversal.
    * Alternative to Iterator
    * Hierarchy members help you traverse themselves.

The visitor is a pattern where a component (visitor) is allowed to traverse the entire hirarchy of types. Implemented by propagating a single Accept() method throughout the entire hierarchy.

### In Summary
* We take every single element in the hierarchy and propagate an *Accept(v *Visitor)* method throughout the entire hierarchy.
* Create a visitor with *VisitFoo(f Foo)*, *VisitBar(b Bar)*,... for each element in the hierarchy.
* Each Accept() simply calls Visitor.VisitX(self). This is useful for both traversal as well as any kind of other concern where we need to go through a set of related elements and get some information about them.