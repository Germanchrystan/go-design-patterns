# Strategy Design Pattern
Many algorithms can be decomposed into higher and lower level parts. What this means is that we can define an algorithm in terms of the general strategy, as well as the implementation detail.

Imagine the preparation of a hot beverage. You can decompose this into the process of the general making of a hot beverage, where you just boil the water and pour it into a cup. Then there are the specifics, like, for example, if you are making a cup of tea, you have to put the tea bag into the water. If you are making coffee, you have to grind the coffee and put it into water, maybe add cream, etc.
You can see there is a separation between the high level procedure, that happens in all cases, and the low level implementation details that you have to provide.
The high level algorithm can be reused in different specific cases. So you can have different beverage specific strategies.

The stretegy pattern separates an algorithm into its 'skeleton' and concrete implementation steps, which can be varied at run-time.

### In Summary
- Define an alogrithm at a high level
- Define the interface you expect each strategy to follow.
- Support the injection of the strategy into the high-level algorithm.