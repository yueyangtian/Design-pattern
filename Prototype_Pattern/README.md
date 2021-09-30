## 原型模式

#### 模式动机

* 原型模式是创建型模式的一种，其特点在于通过“复制”一个已经存在的实例来返回新的实例,而不是新建实例。被复制的实例就是我们所称的“原型”，这个原型是可定制的。原型模式多用于创建复杂的或者耗时的实例，因为这种情况下，复制一个已经存在的实例使程序运行更高效；或者创建值相等，只是命名不一样的同类数据。

#### 模式结构

* 计算机操作系统中的cp 操作就是一个可以利用原型模式的经典案例，对于每个inode有文件属性或者文件夹属性，当执行clone操作时对于不同的inode的属性会有不同的克隆操作

![Class Diagram](http://www.plantuml.com/plantuml/proxy?src=https://raw.githubusercontent.com/yueyangtian/Design-pattern/master/UML/prototype.puml)
