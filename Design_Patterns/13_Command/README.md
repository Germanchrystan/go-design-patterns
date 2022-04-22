# Command Design Pattern
When we write ordinary code statements, this statements are perishable. Suppose you do a field assignement, for example, yo assign the field age to the value of 20. You cannot by default undo that operation because there is no record of it. You also cannot serialize the sequence of actions or calls.
What we want is an object that represents an operation. For example, you might have a command which says that a construct 'person' should change its age to value 22. In this case, you would:
- Define a structure where you would say what construct has to change (person).
- What needs to be changed in this construct (age).
- What value should be changed to (22). 

Tipically, any kind of user interface is done using commands. Commands also allow us to do multilevel undo and redo, by saving every command that was sent to some system, and then allowing for those commands to roll back and forth.
We can also group commands together and save a whole sequence of commands. These are typically called macros in GUI applications.

A Command is an object which represents an instruction to perform a particular action. Contains all the information necessary for the action to be taken. 
A command can be processed by the object that the command operates on, or processed by itself. It can basically have a call method to apply its commands or alternatively, there could also be a command processor, where all the commands are sent.

### In Summary
- Encapsulate all details of an operation in a separate object.
- Define functions for applying the command(either in the command itself, or elsewhere).
- Optionally define instructions for undoing the command.
- Can create composite commands (macros).