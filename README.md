<p align="center">
  <img src="https://helight.info/wp-content/uploads/2018/03/timg-267x300.png" height="300">
  <h1 align="center">
    Go Patterns
    <br>
  </h1>
</p>

A collection of design & application patterns for Go language.

## Creational Patterns

| Pattern | Description | Status |
|:-------:|:----------- |:------:|
| [Factory](/creational/01_factory/)|Provides an interface for creating families of releated objects | ✔ |
| [Abstract Factory](/creational/02_abstractfactory/) | Provides an interface for creating families of releated objects | ✔ |
| [Singleton](/creational/03_singleton) | Restricts instantiation of a type to one object | ✔ |
| [Builder](/creational/04_builder/) | Builds a complex object using simple objects | ✔ |
| [Prototype](/creational/05_prototype) | Instantiates and maintains a group of objects instances of the same type | ✔ |


## Structural Patterns

| Pattern | Description | Status |
|:-------:|:----------- |:------:|
| [Adapter](/structural/06_adapter) |  | ✘ |
| [Bridge](/structural/07_bridge) | Decouples an interface from its implementation so that the two can vary independently | ✘ |
| [Filter](/structural/08_filter) |  | ✘ |
| [Composite](/structural/09_composite) | Encapsulates and provides access to a number of different objects | ✘ |
| [Decorator](/structural/10_decorator) | Adds behavior to an object, statically or dynamically | ✔ |
| [Facade](/structural/11_facade) | Uses one type as an API to a number of others | ✘ |
| [Flyweight](/structural/12_flyweight) | Reuses existing instances of objects with similar/identical state to minimize resource usage | ✘ |
| [Proxy](/structural/13_proxy) | Provides a surrogate for an object to control it's actions | ✔ |

## Behavioral Patterns

| Pattern | Description | Status |
|:-------:|:----------- |:------:|
| [Chain of Responsibility](/behavioral/chain_of_responsibility.md) | Avoids coupling a sender to receiver by giving more than object a chance to handle the request | ✘ |
| [Command](/behavioral/command.md) | Bundles a command and arguments to call later | ✘ |
| [Mediator](/behavioral/mediator.md) | Connects objects and acts as a proxy | ✘ |
| [Memento](/behavioral/memento.md) | Generate an opaque token that can be used to go back to a previous state | ✘ |
| [Observer](/behavioral/observer.md) | Provide a callback for notification of events/changes to data | ✘ |
| [Registry](/behavioral/registry.md) | Keep track of all subclasses of a given class | ✘ |
| [State](/behavioral/state.md) | Encapsulates varying behavior for the same object based on its internal state | ✘ |
| [Strategy](/behavioral/strategy.md) | Enables an algorithm's behavior to be selected at runtime | ✘ |
| [Template](/behavioral/template.md) | Defines a skeleton class which defers some methods to subclasses | ✘ |
| [Visitor](/behavioral/visitor.md) | Separates an algorithm from an object on which it operates | ✘ |

### refer:
	http://www.runoob.com/design-pattern/
