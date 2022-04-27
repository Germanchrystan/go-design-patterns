# Observer Design Pattern
Sometimes we need to be informed when a certain value in an object is changed, or if the object does something. Maybe we want to be informed when an external event occurs and we want out system to be able to handle it somehow. There are different ways of handling these situations.

Suppose we want to detect that an object's field has just changed. Checking the field every hundred milliseconds would not be practical. Instead, what we want is the object to tell us about the change, by generating events for every change. We want to listen to and be notified by these events.

There are two parts in this process then: observable and observer. The observer design pattern involves both participants.

An observer is an object that wishes to be informed about events happening in the system. The entity generating the events is an observable.

### In Summary
- By definition it is an intrusive approach, which means that to make an object observable, we would have to change that object, unfortunately.
- Must provide a way of clients to subscribe.
- Event data sent from observable to all observers.
- Data represented as interface{}, to make data sent as generic aspossible. We let the different components define what kind of data they send.
- Unsubscription is possible.