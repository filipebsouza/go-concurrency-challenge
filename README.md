# Concurrency challenge

This exercise was made for fixed the `golang` concurrency tools with a practical scenario.

# Burger House

The context is a restaurant, actually a Burger House that make `Burgers`, `Fries` and `Beverages`. These products have
different preparation times, but the `Order` just can be delivery after all products are done.

# Requirements

### Burger

Preparation process:

Default time: 02 min

| Point       | Preparation time            |
|-------------|-----------------------------|
| Blue rare   | Default time                |
| Rare        | Default time + 30 sec       |
| Medium rare | Default time + 1 min        |
| Medium      | Default time + 1 min 30 sec |
| Medium well | Default time + 2 min        |
| Well done   | Default time + 2 min 30 sec |

| Bacon         | Preparation time |
|---------------|------------------|
| Without bacon | Nothing          |
| Streaky       | 2 min            |
| Canadian      | 3 min            |
| Crumbs        | 4 min            |

| Salad                | Preparation time |
|----------------------|------------------|
| Without salad        | Nothing          |
| Lettuce              | 1 min            |
| Lettuce and tomatoes | 1 min 10 sec     |
| Creamy coleslaw      | 2 min            |
| Crispy fried cabbage | 2 min            |

**Example:** Well done Burger (4 min 30 sec) with Canadian bacon (3 min) and Lettuce salad (1 min). Total preparation
time: 8 min 30 sec

### Beverage

Preparation process:

Default time: 15 sec

| Kind         | Preparation time      |
|--------------|-----------------------|
| Coke         | Default time          |
| Soda         | Default time          |
| Orange       | Default time          |
| Grape        | Default time          |
| Iced tea     | Default time + 10 sec |
| Orange juice | Default time + 15 sec |

**Example:** Iced tea (25 sec). Total preparation: 25 sec.

### Fries

Preparation process:

Default time: 8 min

| Kind         | Preparation time     |
|--------------|----------------------|
| Normal       | Default time         |
| Spice        | Default time + 2 min |
| Rustic fries | Default time + 4 min |

**Example:** Rustic fries (12 min). Total preparation: 12 min.

### Order

An Order contain one or more products.
An Order starts with the state `Requested`, when the products start the preparation become `Making` and when done become `Ready for delivery`.
