# go 是如何实现面向对象编程的？

!!! quote
    The object-oriented model makes it easy to build up programs by accretion. What this often means, in practice, is that it provides a structured way to write spaghetti code. - Paul Graham

## 面向对象编程

[面向对象编程(OOP)](https://en.wikipedia.org/wiki/Object-oriented_programming) 应该是近几十年最重要的编程范式之一，流行的编程语言
Java/C++/Python 等都支持 OOP。如果你使用过 Java/Python 的 OOP，应该很熟悉一些概念，比如类，对象(实例)，抽象，封装，继承，多态等。

过程式编程中我们一般通过封装成函数来进行逻辑复用。 OOP 中一般我们会根据业务进行抽象，把现实中的实体抽象成一个类(class)，
类包含数据(data)和操作数据的方法(method)，我们可以创建一个类的实例(instance)或者也叫对象(object)来对其进行操作，不同
编程语言中都对 OOP 有类似的支持。

但是初学 go 的童鞋一般会对 go 的面向对象编程感到很新奇，它和常见的编程语言 Python/Java 实现的方式差别比较大。比如 go
里边没有 class 的概念，go 通过 struct 实现自定义类型，而且 go 的 struct 不支持继承，只支持组合等。一开始笔者也感觉 go
这种 oop 方式会很弱，但是写多了你会发现，go 实现 oop 的方式简单又够用。
