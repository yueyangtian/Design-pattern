## 抽象工厂模式
#### 模式动机
在工厂方法模式中具体工厂负责生产具体的产品，每一个具体工厂对应一种具体产品，工厂方法也具有唯一性，一般情况下，一个具体工厂中只有一个工厂方法或者一组重载的工厂方法。但是有时候我们需要一个工厂可以提供多个产品对象，而不是单一的产品对象
#### 模式结构
如图抽象工厂模式本质是多维度的工厂模式，并通过增加对象声明时代码量的代价换摆脱了简单工厂的产品创建的状态机逻辑
![Class Diagram](http://www.plantuml.com/plantuml/proxy?src=https://raw.githubusercontent.com/yueyangtian/Design-pattern/master/UML/abstract_factory.puml)