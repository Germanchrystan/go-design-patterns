# Prototype
A partially or fully initialized object that can be copied (cloned) and make use of.

- Complicated objects aren't designed from scratch. Manufacturers reiterate existing designs.
- An existing (partially or fully constructed) design is a prototype.
- We make a copy of the prototype and customize it. Requires 'deep copy' support.
- We make the cloning convenient (e.g., via a factory).

Steps to implement a prototype: 
- Partially construct an object and store it somewhere.
- Deep copy the prototype.
- Customize the resulting instance.
- A prototype factory provides a convenient API for using prototype.