## 桥接模式

#### 模式动机

* 所谓的代理者是指一个类别可以作为其它东西的接口。代理者可以作任何东西的接口：网络连接、存储器中的大对象、文件或其它昂贵或无法复制的资源。C++中的智能指针以及STL中的迭代器的结构思想都借鉴了这种模式，他们都实现了指针访问的接口功能，并且都依赖于真实访问对象


#### 模式结构

* 如图所示假设不同的电脑品牌的开机欢迎语输出与电脑的品牌和操作系统有关，这时我们可以将操作系统的开机欢迎语通过电脑品牌的开机欢迎语打印方式进行桥接

![Class Diagram](http://www.plantuml.com/plantuml/proxy?src=https://raw.githubusercontent.com/yueyangtian/Design-pattern/master/UML/proxy.puml)