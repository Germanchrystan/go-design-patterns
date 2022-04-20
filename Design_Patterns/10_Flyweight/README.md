# Flyweight
Space optimization technique to avoid redundancy when storing data. 
Imagine we have a backend application that manages user registration and login. It is possible we would have lots of users with the same first and last names. There is no sense in storing the same names over and over again. We could store a list of names and references to them (indices, pointers, etc.).

Other example could be if we had a text formatting software. We don't want each character in the text to have formatting character information (bold or italic), but instead we should operate on ranges (line number, start/end positions).

Flyweight is a space optimization technique that lets us use less memory by storing externally the data associated with similar objects.

In summary, flyweight is a mechanism in which we store common data externally, in order to not store it at the position of the structure that need this data. By specifying an index or a pointer into the external data store, objects can share this data and access it where neccesary. We can define the idea of 'ranges' on homogeneous collections and store data related to those ranges.