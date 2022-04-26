# Interpreter Design Pattern
Interpreters are all around us. Even now, in this very room

Interpreters are used everywhere when it comes to the structure and interpretation of computer programs. 
We recommend Structure and Interpretation of Computer Programs by Harold Abelson and Gerald Jay Sussman for a greater understanding.

Essentially the idea is that the interpreter design pattern handles the situation where we need to process textual input. There are so many different situations where that is required. You could need to turn text into some sort of data structure. Tipically what we call this data structure, especially if we are parsing programming languages, is abstract syntax tree. The tree can subsequently be traversed using, for example, the visitor pattern, and thewn we can transform it, compile it, etc.

Examples: 
- Programming language compilers, interpreters and IDEs. They work by interpreting data. Behind the scenes, every single IDE has a parser in their lexar, for actually understanding the code and perform analysis upon it. 
- Text formats like HTML, XML, JSON
- Numeric expressions that need to be evaluated.
- Regular expressions.

Turning strings into linked structures is a complicated process. In the case of programming languages, it gets particularly difficult. 

An Interpreter is a component that processes structured text data. DOes so by turning it into separate lexical tokens (lexing) and then interpreting sequences of said tokens (parsing).

### In Summary
Barring simple cases, an interpreter acts in two stages.
- Lexing turns text into a set of tokens.
- Parsing turns these tokens into an Abstract Syntax Tree, or some sort of aggregated or linked structure.
That way, this structure can be traversed.