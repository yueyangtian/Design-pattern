## 享元模式

#### 模式动机

* 主要用于减少创建对象的数量，以减少内存占用和提高性能。这种类型的设计模式属于结构型模式，它提供了减少对象数量从而改善应用所需的对象结构的方式。

#### 模式结构

* 通过Flyweight类获取ColorCricle享元，通过Color工厂类在享元类中创建不同属性的colorcircle，在flyweight中通过getCricle方法共享相同属性的colorcirle对象

![Class Diagram](http://www.plantuml.com/plantuml/proxy?src=https://raw.githubusercontent.com/yueyangtian/Design-pattern/master/UML/flyweight.puml)