# Proxy Design Pattern
Imagine you write this piece of code 

```go
    foo.Bar()
```
There are some assumptions behind the scenes here. This assumes that foo is in the same process as Bar().
But does it really have to be the case? What if, later on, you want to put all Foo-related operations into a separate process? So instead of calling in the process that you are currently in, you want the invocation of foo.Bar() to happen in a different process or maybe on a different machine, somewhere on the network. Can you avoid changing your code? That is where the proxy pattern comes to the rescue.

The idea of the proxy design pattern is that we provide the same interface, but the behavior is entirely different. So we would still have a variable called foo and it would still have a method called Bar. But this method would not call something in process, it would call something somewhre else. This is called a communication proxy. Other types of proxies that we would encounter are logging, virtual, guarding, etc.

A proxy is a type that functions as an interface to a particular resource. That resource may be remote, expensive to construct, or may require logging, or some other added functionality.

### Proxy vs. Decorator
One thing worth mentioning is the difference between the proxy and the decorator. They are similar, but they serve different proposes.

The proxy design pattern tries to provide an identical interface whenever it is possible. It tries to provide an identical interface to whatever it is controlling, to whatever resource it is using behind the scenes. Whereas the decorator provides an enhanced interface, with more fields and more methods.

The decorator tipically aggregates (or has pointers to) what it is decorating. The proxy doesn't have to use the underlying object at all. 

The proxy might not even be working with a materialized object. The proxy could still be usable, even if the object for which the proxy works doesn't even exist.

### In Summary
- A proxy has the same interface as the underlying object
- To create a proxy, simply replicate the existing interface of an object
- Add relevant functionality to the redefined methods.