## 中介者模式

#### 模式动机

* 在软件工程领域，中介者模式定义了一个中介者对象，该对象封装了系统中对象间的交互方式。 由于它可以在运行时改变程序的行为，这种模式是一种行为型模式。通常程序由大量的类组成，这些类中包含程序的逻辑和运算。 然而，当开发者将更多的类加入到程序中之后，类间交互关系可能变得更为复杂，这会使得代码变得更加难以阅读和维护，尤其是在重构的时候。 此外，程序将会变得难以修改，因为对其所做的任何修改都有可能影响到其它几个类中的代码。在中介者模式中，对象间的通信过程被封装在一个中介者（调解人）对象之中。 对象之间不再直接交互，而是通过调解人进行交互。 这么做可以减少可交互对象间的依赖，从而降低耦合。

#### 模式结构

* 我们通过聊天室实例来演示中介者模式。实例中，多个用户可以向聊天室发送消息，聊天室向所有的用户显示消息。我们将创建两个类 ChatRoom 和 User。User 对象使用 ChatRoom 方法来分享他们的消息。

![Class Diagram](http://www.plantuml.com/plantuml/proxy?src=https://raw.githubusercontent.com/yueyangtian/Design-pattern/master/UML/mediator.puml)