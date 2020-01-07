# go design patterns

Golang implements design patterns.

## Creational pattern

### Singleton pattern



### Structual design pattern

#### proxy

The proxy pattern wraps an object to hide some of its characteristics.
These characteristics could be the fact that is is a remote object(`remote proxy`),a very heavy object such as a very big image or
the dump of a terabyte database(`virtual proxy`), or a restrictes access object(`protection proxy`).

##### Objectives

- Hide an object behind the proxy so the features can be hidden,restricted and so on
- Provide a new abstraction layer that is easy to work with,and can be changed easily

#### Decorator

Decrator pattern is useful when woring with legacy code. 

##### Objectives

When you think about extending legacy code without the risk of breaking something,you should think of the decrator first.

Like in a Swiss knife, you have a base type (the frame of the knife),and from there yu unfold its functionalities.

##### Acceptance criteria 验收标准

- We must have the main interface that all decorators will implement. This interface will be called `IngredientAdd`,and it will have the `AddIngredient()` string method.
- We must have a core `PizzDecorator` type that we will add infredients to.
- We must have an ingredient "onion" implementing the same `IngredientAdd` interface that will add the string `onion` to the returned pizza.
- ...