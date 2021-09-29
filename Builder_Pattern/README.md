## 建造者模式

#### 模式动机

* 无论是在现实世界中还是在软件系统中，都存在一些复杂的对象，它们拥有多个组成部分，而对于大多数用户而言，无须知道这些部件的装配细节，也几乎不会使用单独某个部件，而是使用一辆完整的汽车，可以通过建造者模式对其进行设计与描述，建造者模式可以将部件和其组装过程分开，一步一步创建一个复杂的对象。用户只需要指定复杂对象的类型就可以得到该对象，而无须知道其内部的具体构造细节

#### 模式结构

* 如图通过CarBuilder建造者对汽车颜色以及时速属性进行构建
![Class Diagram](http://www.plantuml.com/plantuml/proxy?src=https://raw.githubusercontent.com/yueyangtian/Design-pattern/master/UML/builder.puml)
