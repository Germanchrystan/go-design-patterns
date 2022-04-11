# Bridge
Bridge prevents a particular problem, which is tipically called a 'Cartesian product' complexity explosion.

Example: We want to have thread scheduler. A thread scheduler can be preemptive or corporative, and can run on either Windows or Linux. So, if we tackle this problem straight on, we end up with a two by two scenerario.
    - Windows Preemptive Scheduler.
    - Windows Corporative Scheduler.
    - Unix Preemptive Scheduler.
    - Unix Corporative Scheduler.

So, we would have four different structures. If a third operating system was introduced, as well a different type of scheduler, we would end up with nine variations. This can become unmanageable.

So this is why the Bridge design pattern is actually for. It tries to avoid this complexity explosion.
What we would do in this example is to have a top level construct named "ThreadScheduler", which in turn can be aggregated inside a preemptive or cooperative type. On the other hand, we could also have a reference or a pointer to include a platform scheduler.

So instead of having one huge tree, we could have a couple of flat trees or lists.

The Bridge design pattern is a mechanism that decouples an interface (hierarchy) or abstraction from an implementation (hierarchy).

It is a stronger form of encapsulation.
