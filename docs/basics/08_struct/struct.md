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
我们就需要使用指针接收者。如果你不好判断使用指针还是值接收者，推荐你使用指针接收者。

!!! warning
    注意代码里的 NOTE 注释，go 里提供了简化指针访问成员的方式，比如我们没有使用 `(*a).petName` 而是直接使用的
    `a.petName = petName`。这里并不是语法错误，而是 go 提供的一个好用的语法糖，让我们可以直接用这种方式访问成员。

## 构造函数如何实现？

上文我们是通过初始化一个 Animal 结构体的方式创建了一个 Animal "对象"，go
里并没有像其他语言那样提供构造函数的方式来创建一个对象(是不是很无趣，对 go 就是这么吝啬)。
我们知道 go 里边创建一个空结构体的时候，不同类型的成员被赋值成其类型的『零值』，比如对于 Animal 里的
Name(string)是空字符串, Age(int) 是 0。

那如果我们想创建一个 Animal 的时候根据传入的参数来初始化呢？go 里边虽然没有直接提供构造函数，但是一般我们是通过定义一些
NewXXX 开头的函数来实现构造函数的功能的，比如我们可以定义一个 NewAnimal 函数：

``` go hl_lines="15 16 17 18 19 20 21"
package main

import "fmt"

type Animal struct {
	Name    string
	Age     int
	petName string
}

func (a Animal) Sleep() {
	fmt.Printf("%s is sleeping\n", a.Name)
}

func NewAnimal(name string, age int) *Animal {
	a := Animal{
		Name: name,
		Age:  age,
	}
	return &a
}

func main() {
	a := NewAnimal("cat", 3)
	fmt.Println(a)
}
```

这样我们就实现了一个类似构造函数的功能，当然你也可以根据不同的需求来定义多个构造函数。

## 组合 vs 继承

上文学习了如何定义 go 的 "对象"，我们给 struct 加入了数据成员和方法，还实现了构造函数，看起来稍微有点面向对象编程的意思了。
OOP 中还有一个重要的概念就是继承，通过继承实现了 is-a 的类关系，可以很好地进行代码复用。但是 go 可能又要让你失望了，你会
发现 go 并不直接支持 struct 之间的继承。

那如果我们想实现类似继承的功能该怎么办呢？其实 go 也有类似的解决方案，不过 go 使用的不是继承而是组合，go 作者推崇的
思想是『组合优于继承』。go 提供了结构体的嵌入(embedding)用来实现代码复用，比如如果我们想定义一个 Dog 结构体，Dog 也是一个
Animal，我们想复用 Animal 里的成员，可以在 Dog struct 里嵌入一个 Animal:

```go
type Dog struct {
	Animal // embedding

	Color string
}

func main() {
	d := Dog{}
	d.Name = "dog"
	d.Sleep()
}
```

你会发现在 Dog 里嵌入了 Animal 以后，我们就可以使用 Animal 的成员和方法了，从而实现了代码复用，是不是实现起来很简单。
我们还可以重写 Dog 自己的 Sleep 方法，来覆盖掉 Animal 的 Sleep 方法，给 Dog 增加一个方法:

```go hl_lines="7 8 9"
type Dog struct {
	Animal // embedding

	Color string
}

func (d Dog) Sleep() {
	fmt.Println("Dog method Sleep")
}

func main() {
	d := Dog{}
	d.Name = "dog"
	d.Sleep() // 输出的是 Dog 的 Sleep 方法而不是 Animal 的
}
```

类似的，如果嵌入的 struct 里的成员名字和当前 struct 同名冲突了，go 会优先使用当前 struct 的成员。
到这里我们就大概学习了 go 使用 struct 来实现 OOP 的方式，可以看得出和常用的编程语言 Java/C++/Python 等还是有不少的区别的。
总得来说，go 的设计就是大道至简，没有其他语言那么多复杂的概念和语法糖，甚至让人感觉比较『简陋』。但是用多了你会发现，go
的这种设计精简并且够用，并且大大简化了代码的学习和上手成本。

## 多态

到这里我们还有一个 OOP 中重要的概念没有介绍，就是多态的概念。简单的说，多态就是同一个接口，对于不同的实例执行不同的操作。
下一章我们将介绍下 go 的接口(interface)，以及如何在 go 中实现多态。

## 练习

- 之前我们学过 go 的 map，但是 go 里边没有直接提供一个 set，请你使用 struct 封装一个 Set，并且提供 Add/Delete/Exist 方法
- 通过使用嵌入实现一个 Cat struct，加入一个数据成员叫做 Height，并且给你的 Cat 加上 Eat 方法。

## 参考

- [Effective Go pointers vs. Values](https://golang.org/doc/effective_go.html#pointers_vs_values)
- [Object Oriented Inheritance in Go](https://hackthology.com/object-oriented-inheritance-in-go.html)
- [Object-oriented programming without inheritance](https://yourbasic.org/golang/inheritance-object-oriented/)
