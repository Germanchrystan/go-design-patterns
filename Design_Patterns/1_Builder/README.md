# BUILDER DESIGN PATTERN        

Some objects are simple and can be created in a single constructor call. In other situations, objects require a lot of ceremony to create.
Having a factory function with 10 arguments is not productive. It means forcing the user to make lots of decisions within a single expression, and that is never a good thing.

We want to somehow make the construction process a sort of a multi-stage process. That way, we can construct an object piece-wise, as opposed to trying to do everything in a single factory call.

The builder pattern is basically about providing some sort of API for constructing an object step by step, as opposed to trying to construct it all at once. 

What is often useful is to create a fluent builder. The fluent builder methods return the receiver's pointer, which allows chaining of building methods one after another. 

Different facets of an object can also be constructed with several different builders. If you have a really complicated object that has many different aspects, those could be separated between separate builders with separate concerns.