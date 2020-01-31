# go struct

!!! quote
    I think a lot of new programmers like to use advanced data structures and advanced language features as a way of demonstrating their ability. I call it the lion-tamer syndrome. Such demonstrations are impressive, but unless they actually translate into real wins for the project, avoid them. - Glyn Williams

## go 支持 OOP 么？

细心的你可能发现了，go 连 class 关键字都没有，如何支持面向对象编程呢？流行的编程语言一般都支持定义一个类，类里边有数据(data)和方法(method)。
go 虽然没有提供 class 关键词，但是提供了 struct 用来定义自己的类型，struct 里可以放入需要的数据成员，并且可以给自定义
struct 增加方法。

## 使用 struct 自定义类型

我们开始使用 go 的 struct 来看一下 go 里边是如何实现类似其他编程语言的面向对象编程的。在 go 语言里，我们使用 struct
来定义一个新的类型，这和使用 class 非常像，比如这里我们定义一个简单的动物结构（类），包含 Name 和 Age 成员：

```go
package main

import "fmt"

type Animal struct {
	Name string
	Age  int
}

func main() {
	a := Animal{Name: "dog", Age: 3}
	fmt.Println(a, a.Name, a.Age)
}
```

ok，然后还可以给 struct 定义方法，比如动物都需要睡觉，所以我们给它添加一个方法叫做 Sleep：

``` go hl_lines="10 11 12 17"
package main

import "fmt"

type Animal struct {
	Name string
	Age  int
}

func (a Animal) Sleep() {
	fmt.Printf("%s is sleeping", a.Name)
}

func main() {
	a := Animal{Name: "dog", Age: 3}
	fmt.Println(a, a.Name, a.Age)
	a.Sleep()
}
```

这样定义既有数据又有了方法，是不是和类比较像了，这就是在 go 中使用 OOP 的方式。当然了 OOP 还远不止这些，比如传统的
面向对象编程还有访问控制，构造函数，继承，多态等概念，我们会看一下它们是如何在 go 里边实现的，go 的方式和 Python/Java
等实现还是有挺大区别的。

## 访问控制

在 Java 和 C++ 中，对于类的成员有着比较严格的访问控制，比如对成员有 public/private 等关键词用来声明它的访问权限。但是像是
Python 的实现就没有那么严格，Python 里边是通过命名的方式来约定的，比如私有方法和成员一般是用下划线开头，告诉调用者这个是
类的私有方法和成员，你不应该直接使用它们，而是使用暴露出去的公有方法。但这只是一个『君子协定』，如果类的设计者没有
设计完善，让你非要去访问下划线开头的『私有方法』，其实 python 也不会禁止。

Go 也有类似的访问控制，不过是通过数据和方法的命名**首字母大小写**决定的。在 Go 的包(package)中，只有首字母大写的才能被其他包
导入使用，小写开头的则不行。所以一般结构体的私有数据成员和方法，我们使用小写开头，而公有的数据成员和方法，我们使用大写开头就好了。

我们现在给 Animal 加上一个私有的数据成员叫做 petName 表示动物的小名，同时新增一个方法叫做 SetPetName 用来设置它：

```go hl_lines="8 15 16 17 25"
package main

import "fmt"

type Animal struct {
	Name    string
	Age     int
	petName string
}

func (a Animal) Sleep() {
	fmt.Printf("%s is sleeping", a.Name)
}

func (a Animal) SetPetName(petName string) {
	a.petName = petName
}

func main() {
	a := Animal{Name: "dog", Age: 3}
	fmt.Println(a, a.Name, a.Age)
	a.Sleep()

	a.SetPetName("little dog")
	fmt.Println(a.petName) // 为什么没有设置成功？
}
```

打印上边代码看下结果？符合你的预期么？如果没有正确设置 petName 为什么呢？

## 指针接收者(pointer receiver) vs 值接收者(value receiver)

在函数章节我们讲到过函数的参数传递都是通过值进行传递的，也就是会复制参数的值，如果我们想要修改传入的值，就需要传递一个指针，
这就是为什么上边的 SetPetName 方法没有起作用的原因。如果我们想要修改一个结构体，就需要传入一个结构体指针作为接收者：
(注意这里依然是值拷贝，不过拷贝的是指针的值而不是整个结构体的值，通过指针就可以修改对应的结构体)

``` go hl_lines="15 16 17 18 19"
package main

import "fmt"

type Animal struct {
	Name    string
	Age     int
	petName string
}

func (a Animal) Sleep() {
	fmt.Printf("%s is sleeping", a.Name)
}

func (a *Animal) SetPetName(petName string) {
	a.petName = petName // NOTE: 这里的 a 是一个指针
	// NOTE: 以下这种方式也是可以的，go 如果碰到指针会自动帮我们处理，所以使用起来更方便
	// (*a).petName = petName
}

func main() {
	aPtr := &Animal{Name: "dog", Age: 3}
	aPtr.SetPetName("little dog")
	fmt.Println(aPtr.petName) // 是不是可以设置成功了
}
```

运行代码你就会发现我们成功修改了 Animal 的 petName 成员了。一般如果必须需要修改结构体，或者结构体数据成员比较多（减少复制成本），
我们就需要使用指针接收者。

注意代码里的 NOTE 注释的细节问题，go 里提供了简化指针访问成员的方式，比如我们没有使用 `(*a).petName` 而是直接使用的
`a.petName = petName`。这里并不是语法错误，而是 go 提供的一个好用的语法糖，让我们可以直接用这种方式访问成员。

## 构造函数如何实现？

## 组合 vs 继承

## 多态

## 练习

- 实现 set

## 参考

- [Effective Go pointers vs. Values](https://golang.org/doc/effective_go.html#pointers_vs_values)
- [Object Oriented Inheritance in Go](https://hackthology.com/object-oriented-inheritance-in-go.html)
- [Object-oriented programming without inheritance](https://yourbasic.org/golang/inheritance-object-oriented/)
