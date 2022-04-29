# Template Method Design Pattern

* Algorithms can be decomposed into common parts plus specifics.
* Strategy pattern does this through composition.
    * High-level algorithm uses an interface.
    * Concrete implementations implement the interface.
    * We keep a pointer to the interface; provide concrete implementations.
* Template Method performs a similar operation, but
    * It is typically just a function, not struct with a reference to the implementation.
    * Can still use interfaces (just like Strategy); or
    * Can be functional (take several functions as parameters), and in this case we can do without structs or interfaces, we can just work with functions and nothing but functions. This is the functional template method approach.

A template method is a skeleton algorithm defined in a function. Function can either use an interface (like Strategy) or can take several functions as arguments.

### In Summary
* Very similar to Strategy.
* Typical implementation:
    * Define an interface with common operations.
    * Make use of those operations inside a function.
* Alternative functional approach:
    * Make a function that takes several functions
    * Can pass in functions that capture local state.
    * No need for either structs or interfaces.