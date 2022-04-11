# Adapter
Electrical devices have different power (interface) requirements. You might have different voltage requirements or different sockets types. If I have a laptop with a plug, I cannot modify the plug to suddenly support every single possible interface. This is where adapters come in handy. An adapter is a device designed to give us the interface we require.

In software engineering, we have this idea of an adapter as a construct which adapts some existing interface X to conform to the required interface Y.

Implementing an Adapter is easy:
- Determine the API you have and the API you need.
- Create a component which aggregates(has pointer to, ...) the adaptee.
- Intermediate representations can pile up: use caching and other optimizations: make sure that the amount of temporary data generated as the adapter is being executed is manageable and it doesn't go out of bounds.