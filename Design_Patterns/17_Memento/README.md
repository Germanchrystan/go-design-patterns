# Memento Design Pattern
Keep a memento of an object's state to return to that state

As we have objects in a system, those objects go through changes.
Now, there are different ways of actually navigating those changes. One way is to record every change by using, for instance, the command design pattern, which could have an undo function. This is part of something called Command Query Responsibility Segregation, an approach which sometimes is also coupled with event sourcing. That is something we would typically discuss as part of enterprise patterns, as opposed to conventional design patterns.

Another approach is to simply save snapshots of the system, so whenever we need to, we can returned to a previously saved snapshot.

The Memento is a token representing the system state. It lets us roll back to the state when the token was generated. It may or may not directly expose state information. 

### Memento vs. Flyweight
- Both patterns provide a 'token' clients can hold on to.
- The memento is used only to be fed back into the system. It has no public or mutable state, nor any methods.
- A Flyweight is similar to an ordinary reference to object. It can mutate its state, and provide additional functionality (fields and methods).

### In Summary
- Mementos are used to roll back states arbitrarily.
- A memento is simply a token/handle with typically no methods of its own. If it has methods, we should try to hide them as much as possible. Changing a memento is un-idiomatic, it should be a read-only object.
- A memento is not required to expose directly the state(s) to which it reverts the system. It can be problematic if it does, because if it exposes some state, then somebody could modify it and then return the system to the state in which it never was. 
- It can be used to implement undo/redo. This implementation is somewhat clunky, because in order for this to work, we basically have to save every single memento. It is not practical to have saved states of every single point in time. If the system is complicated, we will be taking huge snapshots with lots of data being replicated over and over again for each memento. That would be too much computation and memory traffic to make this approach realistic. Undo and Redo functionality would be better handled by the command design pattern.
