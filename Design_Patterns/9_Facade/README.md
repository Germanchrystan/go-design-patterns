# Facade Design Pattern
We want to balance complexity and presentation/usability.
Think of a house: it has lots of different subsystems. Electrical, plumbing, ventilation, etc. This complex and multilayered internal structure is not perceived by or exposed to the end user. A similar thing happens with software. Sometimes we have a complicated system or indeed a set of subsystems behind the scenes, but we don't want or care to delve into this complexity. We just care about using a flexible API.

Facade is a mechanism that provides a simple, easy to understand/user interface over a large and sophisticated body of code.

We may also whish to optionally expose those internal details so that if we have a power user, they can also manipulate those implementation details. That way, users can escalate to use more complex APIs if they need to.