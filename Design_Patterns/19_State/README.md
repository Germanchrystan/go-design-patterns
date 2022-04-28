# State Design Pattern

Consider an ordinary telephone. What you do with the telephone depends on the state of the phone and the state of the phone line. If the phone is ringing or if you want to make a call, you can pick up the phone. To make a call and talk to somebody the phone must be off the hook. You cannot talk to people while the phone is still on the hook. If you try to call someone, and it is busy, you may leave a message on the answering machine.

The idea is that we have all these changes in state. Changes in state can be explicit or in response to an event (observer pattern). 

The State design pattern is a pattern in which the object's behavior is determined by its state. An object transitions from one state to another (something needs to trigger a transition).

A formalized construct which manages state and transitions is called a state machine.
We also use the term finite state machine when the state machine has a specific starting state and also a specific terminal state, after which the execution of the state machine is finished.

### In Summary
- Given sufficient complexity, it pays to formally define possible states and events/triggers.
- We can define:
--- state entry/exit behaviors.
--- Actions when a particular event causes a transition
--- Guard conditions enabling/disabling a transition.
--- Default action when no transitions are found for an event.