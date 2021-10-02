## 修饰器模式

#### 模式动机

* 又叫部分整体模式，是用于把一组相似的对象当作一个单一的对象。组合模式依据树形结构来组合对象，用来表示部分以及整体层次。这种类型的设计模式属于结构型模式，它创建了对象组的树形结构。

#### 模式结构
* 如果将文件系统从结构上来看的话，对于每一集目录下的所有文件都有统一的文件属性，对于目录来说他与文件对象为组合关系，在此基础上我们模仿文件的查找功能

![Class Diagram](http://www.plantuml.com/plantuml/proxy?src=https://raw.githubusercontent.com/yueyangtian/Design-pattern/master/UML/decorator.puml)