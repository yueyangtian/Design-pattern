## 修饰器模式

#### 模式动机

* 一种动态地往一个类中添加新的行为的设计模式。就功能而言，修饰模式相比生成子类更为灵活，这样可以给某个对象而不是整个类添加一些功能。

#### 模式结构
* 相比于桥接模式，修饰器模式会用被修饰对象初始化实例，并且与被修饰对象为包含关系，并没有对被修饰对象的方法进行重写，可以更加灵活的选择调用修饰器方法或者修饰对象的方法

![Class Diagram](http://www.plantuml.com/plantuml/proxy?src=https://raw.githubusercontent.com/yueyangtian/Design-pattern/master/UML/decorator.puml)