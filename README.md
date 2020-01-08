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


### Behavioral Patterns

We will encapsulate behaviors, for example, algorithms in the Stragy pattern or executions in the command pattern.

#### Strategy design pattern


The Strategy pattern uses different algrithoms to achieve some specific functionality.

These algorithms are hidden behind an interface and they must be interchangeable.

``` go
switch inputType{
    case Text:
    return &TextSquare{

    }
    case Image:
    return &Imagequare{
    
    }
    default:
    return nil,errors.New("No strategy impl yet)
}

```
`TextSquare`, `Imagequare` implements an interface `PrintStrategy`to Print.

在现实当中，字节流的处理就是一种。根据处理的类型，选择具体的处理实现。

### Chain of responsibility design pattern

责任链模式

