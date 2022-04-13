# Decorator Design Pattern
Imagine we want to take and existing object and augment this object with additional functionality.
We don't want to rewrite or alter the object, because we want to stick by the Open Closed Principle.
(open for extension, closed for modification).
We also want to keep enhancements to a particular object, separate from the original object we want to modify
(Separate Responsability Principle). So, we are adhering to two SOLID design principles at once in this particular design pattern.
We want to be able to interact with existing structures, meaning that when we make a decorator over some structure, we also want access to that original structure, with its data and methods.

The solution here is to use embedding. By embedding the decorated object into our new object, the decorator, and then we provide additional functionality.

The decorator is a mechanism that facilitates the addition of behaviors to individual objects through embedding.
It is often used to emulate multiple inheritance.