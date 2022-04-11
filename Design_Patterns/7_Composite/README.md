# Composite Design Pattern
Sometimes, we have a situation where objects use other objects, fields or methods through the idea of embedding.
Composition lets us make compound objects, so we can have objects made up of objects, made up of objects, and so on.
For example, a mathematical expression composed of simple expressions; or a shape group made of several different shapes.

The composite design patterns allows us to treat both single objects and composite objects uniformly, meaning that they would have the same interface. It is a mechanism for treating individual objects (scalar) and compositions of objects in a uniform manner.

In summary:
- Objects can use other objects via composition
- Some composed and singular objects need similar/identical behaviors.
- Composite design pattern lets us treat both types of objects uniformly.
- Iteration supported with the iterator design pattern. (See Iterator chapter)