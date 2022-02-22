# Concurrency challenge

This exercise was made for fixed the `golang` concurrency tools with a practical scenario.

# Burger House

The context is a restaurant, actually a Burger House that make `Burgers`, `Fries` and `Beverages`.
These products have different preparation times, but the `Order` just can be delivery after all products are done.

# Requirements

### Burger

Preparation process:

Default time: 02 min 

| Point      | Preparation time            |
|------------|-----------------------------|
| BlueRare   | Default time                |
| Rare       | Default time + 30 seg       |
| MediumRare | Default time + 1 min        |
| Medium     | Default time + 1 min 30 seg |
| MediumWell | Default time + 2 min        |
| WellDone   | Default time + 2 min 30 seg |



