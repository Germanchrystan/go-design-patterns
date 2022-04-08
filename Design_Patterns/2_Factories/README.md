# Factories
Object creation logic can sometimes become too convoluted. Suppose you have a struct which has lots of lists and maps in it. You want to initialize those, but you don't want to initialize them one by one. 
Every single time that somebody wants to use it, you want to give them default values.
So you create extra constructs to handle this.
Wholesale object creation (non-piecewise, unlike Builder) can be outsourced to:
- A separate function (Factory Function, a.k.a. Constructor)
- That may exist in a separate struct (Factory) for the sake of organization.


A Factory is a component responsible solely for the wholesale (not piecewise) creation of objects.
