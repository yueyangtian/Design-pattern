## 观察者模式

#### 模式动机

* 一个目标对象管理所有相依于它的观察者对象，并且在它本身的状态改变时主动发出通知。这通常透过呼叫各观察者所提供的方法来实现。此种模式通常被用来实时事件处理系统。

#### 模式结构

* 一个文本类，提供输入与回滚的基本操作

![Class Diagram](http://www.plantuml.com/plantuml/proxy?src=https://raw.githubusercontent.com/yueyangtian/Design-pattern/master/UML/observer.puml)