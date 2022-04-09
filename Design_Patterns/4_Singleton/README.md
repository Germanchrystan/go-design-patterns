# Singleton
For some components it only makes sense to have one in the system.
    - Database repository.
    - Object factory.

The construction call can be expensive.
    - We only do it once.
    - We give everyone the same instance.

We want to prevent anyone creating additional copies.

We need to take care of lazy instantiation. We don't want to preemptively load a database if nobody is going to use it.

A singleton is a component that can be instantiated only once.