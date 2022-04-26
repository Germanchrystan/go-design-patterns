# Mediator Design Pattern
Facilitates communication between components

If we have a system where components can go in and out of the system at any time, meaning they can be created or destroyed, it would not make sense to have pointer from one component to another. If we had a chat room application, and a guest in a chat room can leave at any moment, then by having pointers between guests would result in lots of dead pointers for members that left.

The solution to this problem is to have all components refer to some central component that facilitates communication.

The mediator is a component that facilitates communication between other components without them necessarily being awere of each other, or having direct reference or access to each other.

### In summary
- Create the mediator and have each object in the system point to it.
- Mediator engages in bidirectional communitaction with its connected components.
- Mediator has methods the components can call.
- Components have methods the mediator can call.
- Event processing (e.g., Rx) libraries make communication easier to implement.