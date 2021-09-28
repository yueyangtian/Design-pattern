## 工厂模式
#### 模式动机
* 当一个软件系统为了提供一组功能行为一致，表现方式不一致的类型时，就可以使用工厂模式来避免代码的重复
#### 模式结构
* 如图说要对设计不同图形打印的功能，对于用户来说图形的种类是不需要知道的，于是我们可以对图像的打印功能进行抽象，利用多态对图形对象的代码进行复用

![Class Diagram](http://www.plantuml.com/plantuml/proxy?src=https://raw.githubusercontent.com/yueyangtian/Design-pattern/master/UML/factory.puml)
