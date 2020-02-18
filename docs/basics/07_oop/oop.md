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
这种 oop 方式比较简陋，但是写多了你会发现，go 实现 oop 的方式简单又够用。

## 如何自定义类型？

go 支持使用 type 关键词来定义自己的类型，比如我们来定义一个 Enum 类型，go 默认没有提供 enum 类型，我们可以通过 type
自己定义一个枚举类型(在业务代码中经常用到枚举):

```go
package main

import "fmt"

// 自定义一个 Enum类型
type Enum int

const (
	// 这里如果是自增用 iota 更好
	Init    Enum = 0
	Success Enum = 1
	Fail    Enum = 2
)

func main() {
	fmt.Println(Init) // 0
}
```

## 给自定义类型添加方法

上面我们自己定义了一个 Enum 类型，但是它的使用有些局限，比如你可以试试如下代码：

```go
func main() {
	status := 0
	fmt.Println(Init == status) // main.go|18 col 19| invalid operation: Init == status (mismatched types Enum and int)
}
```

你会发现虽然 Enum 是使用 int 定义的，但是你是无法直接进行比较的，go 认为它们是不同的类型。怎么办呢？
你可以使用 int 来进行类型转换，比如使用 `fmt.Println(int(Init) == status)`，这里我们使用另一个种方式，
就是给自定义类型增加方法(method)。

go 允许我们给自定义类型定义方法，所谓的方法其实就是有接收者(receiver)的函数，之前我们已经介绍过函数的定义格式如下：

```go
func functionName(optionalParameters) optionalReturnType {
  body
}
```

方法的定义方式比较类似，只不过多了一个接收者，你可以理解为方法就是有接收者的函数，它的格式如下：

```go
func (r Receiver) functionName(optionalParameters) optionalReturnType {
  body
}
```

比如我们要给 Enum 定一个方法叫做 Int()，它返回  Enum 对应的 int 值，可以这么写:

```go
func (e Enum) Int() int {
	return int(e)
}
```

这样一来就可以直接使用

```go
func main() {
	status := 0
	fmt.Println(Init.Int() == status) // 调用 Init 的 Int 方法返回 int
}
```

一般业务代码里边我们还会给所有状态定义对应的中文或者英文字符串，完整的代码和使用示例如下：

```go
package main

import "fmt"

// 自定义一个 Enum类型
type Enum int

const (
	// 这里如果是自增用 iota 更好
	Init    Enum = 0
	Success Enum = 1
	Fail    Enum = 2

	// 枚举对应的中文
	InitName    = "初始化"
	SuccessName = "成功"
	FailName    = "失败"
)

func (e Enum) Int() int {
	return int(e)
}

func (e Enum) String() string {
	return []string{
		InitName,
		SuccessName,
		FailName,
	}[e]
}

func main() { // 测试一下我们自己定义的 Enum
	status := 0
	fmt.Println(Init.Int() == status)

	status2 := Fail
	fmt.Println(status2.String())
}
```

这样我们就自己定义了一个业务代码常用的枚举类型，通过给自定义类型添加方法，我们可以给类型加入非常多有用的功能。
可以看到方法和普通函数相比，就是多了一个接收者， `func (e Enum) String() string {}`，之后学习 struct 定义方法的时候也是类似的。

除了基本类型，我们还可以自定义一些复杂类型，比如以下一些例子：

```go
func main() {
	// 定义一个 counter 类型
	type Counter map[string]int
	c := Counter{}
	c["word"]++
	fmt.Println(c)

	type Queue []int
	q := make(Queue, 0)
	q = append(q, 1)
	fmt.Println(q)
}
```

可以看到相比使用内置类型，我们自己命名的 Counter/Queue 等含义更加清晰和明确，还能通过增加方法实现更多功能。
除了基于内置类型，我们还可以使用 go 提供的结构体 struct 来定义自己的类型。下一篇文章将介绍下如何使用 struct
来实现面向对象编程。

## 源码延伸

通过看一些 go 的源码，我们可以学习并且模仿 go 的惯用法，比如本文提到的 Enum 类型，在 go 的源码你可以找到类似实现。
以下是 go 的内置的 http server 中关于枚举的实现方式(去掉了注释)：

```go
// src/net/http/server.go
type ConnState int

const (
	StateNew      ConnState = iota
	StateActive
	StateIdle
	StateHijacked
	StateClosed
)

var stateName = map[ConnState]string{
	StateNew:      "new",
	StateActive:   "active",
	StateIdle:     "idle",
	StateHijacked: "hijacked",
	StateClosed:   "closed",
}

func (c ConnState) String() string {
	return stateName[c]
}
```

在你的业务代码中，你也可以使用类似的方式来实现枚举类型。经常在调用一些内置函数的时候跳进去看看，可以让你学到很多 go 的
习惯用法，也推荐你经常看看。

## 小练习

- 请你给 Enum 实现对应的返回英文名称的功能
- 尝试使用自己定义的 Counter 类型来统计一个字符串中不同单词的个数，你可以统计一个文件中的单词个数么？

## 参考

- [Ultimate Visual Guide to Go Enums and iota](https://blog.learngoprogramming.com/golang-const-type-enums-iota-bc4befd096d3)
