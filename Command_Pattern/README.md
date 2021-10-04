## 命令模式

#### 模式动机

* 它尝试以对象来代表实际行动。命令对象可以把行动(action) 及其参数封装起来，于是这些行动可以被重复多次，取消（如果该对象有实现的话），取消后又再重做。


#### 模式结构

* 如图所示我们可以通过开关类表示一个多功能遥控器，他可以开启电视与灯，并且可以记录下操作记录

![Class Diagram](http://www.plantuml.com/plantuml/proxy?src=https://raw.githubusercontent.com/yueyangtian/Design-pattern/master/UML/command.puml)