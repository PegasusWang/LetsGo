# go 接口(interface)

!!! quote
    When I see a bird that walks like a duck and swims like a duck and quacks like a duck, I call that bird a duck.
    –James Whitcomb Riley

## 接口(interface)

如果你使用过 Python/Ruby 之类的动态语言，应该会对[『鸭子类型』](https://baike.baidu.com/item/%E9%B8%AD%E5%AD%90%E7%B1%BB%E5%9E%8B)比较熟悉。

> 当看到一只鸟走起来像鸭子、游泳起来像鸭子、叫起来也像鸭子，那么这只鸟就可以被称为鸭子。

go 通过接口实现了类型安全的鸭子类型，同时又避免了OOP 编程中的继承，我们来学习下 go 的接口(interface)。
go 的接口是一种抽象的自定义类型，没法直接实例化，它声明了一个或者多个方法的签名。如果一个 struct
实现了一个接口定义的所有方法，我们就说这个 struct 实现了这个接口。注意这里的『实现』是隐式的，你不用显示声明某个 struct
实现了哪个接口。

我们来看一个简单的例子，上一章学习 struct 时候，我们定义了一个 Animal，它有一个 Sleep 方法。这里我们定义一个叫做 Sleeper
的接口(go 喜欢用 er 给一个接口作为后缀，比如Reader/Writer)：

```go
// Sleeper 接口声明
type Sleeper interface {
	Sleep() // 声明一个 Sleep() 方法
}
```

然后定义两个 struct，一个猫(Cat)和一个狗(Dog)，并且它们都实了 Sleep 方法，也就是说隐式实现了 Sleeper 接口。

```go
type Dog struct {
	Name string
}

func (d Dog) Sleep() {
	fmt.Printf("Dog %s is sleeping\n", d.Name)
}

type Cat struct {
	Name string
}

func (c Cat) Sleep() {
	fmt.Printf("Cat %s is sleeping\n", c.Name)
}
```

好了，然后我们编写一个函数，不过为了支持多态，函数的参数是一个接口类型而不是具体的 struct 类型。

```go
func AnimalSleep(s Sleeper) {
	s.Sleep()
}
```

完整代码如下：

```go
package main

import (
	"fmt"
)

// Sleeper 接口声明
type Sleeper interface {
	Sleep() // 声明一个 Sleep() 方法
}

type Dog struct {
	Name string
}

func (d Dog) Sleep() {
	fmt.Printf("Dog %s is sleeping\n", d.Name)
}

type Cat struct {
	Name string
}

func (c Cat) Sleep() {
	fmt.Printf("Cat %s is sleeping\n", c.Name)
}

func AnimalSleep(s Sleeper) { // 注意参数是一个 interface
	s.Sleep()
}

func main() {
	var s Sleeper
	dog := Dog{Name: "xiaobai"}
	cat := Cat{Name: "hellokitty"}
	s = dog
	AnimalSleep(s) // 使用 dog 的 Sleep()
	s = cat
	AnimalSleep(s) // 使用 cat 的 Sleep()

  // 创建一个 Sleeper 切片
	sleepList := []Sleeper{Dog{Name: "xiaobai"}, Cat{Name: "kitty"}}
	for _, s := range sleepList {
		s.Sleep()
	}
}
```

ok，到这里我们就用 go 的接口实现了多态，我们先声明了一个接口类型的值，只要实现了这个接口的 struct 变量，都可以赋值给它，
而调用方法的时候，go 会根据实际类型选择使用哪个 struct 的方法。

## 接口嵌入

我们知道 go 的 struct 可以通过嵌入实现代码复用，go 的接口也支持[嵌入](https://golang.org/doc/effective_go.html#embedding)，
来看一个 go 标准库的例子。go 标准库里边定义了 Reader 和 Writer 接口如下：

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
```

只要一个结构体实现了 Read 或者 Write 方法，它就分别实现了 Reader 和 Writer 接口。go 还支持接口嵌套，比如我们可以嵌套这俩
接口声明一个新的接口 ReadWriter:

```go
// ReadWriter is the interface that combines the Reader and Writer interfaces.
type ReadWriter interface {
    Reader
    Writer
}
```

我们也来试一下，刚才声明了 Sleeper 接口，再来声明一个叫做 Eater 的接口。有了睡和吃，我们再组合一下搞一个叫做 LazyAnimal
的接口（只知道吃和睡能不懒么？):

```go hl_lines="16 17 18 19"
package main

import (
	"fmt"
)

// Sleeper 接口声明
type Sleeper interface {
	Sleep() // 声明一个 Sleep() 方法
}

type Eater interface {
	Eat(foodName string) // 声明一个Eat 方法
}

type LazyAnimal interface {
	Sleeper
	Eater
}

type Dog struct {
	Name string
}

func (d Dog) Sleep() {
	fmt.Printf("Dog %s is sleeping\n", d.Name)
}

func (d Dog) Eat(foodName string) {
	fmt.Printf("Dog %s is eating %s\n", d.Name, foodName)
}

type Cat struct {
	Name string
}

func (c Cat) Sleep() {
	fmt.Printf("Cat %s is sleeping\n", c.Name)
}

func (c Cat) Eat(foodName string) {
	fmt.Printf("Cat %s is eating %s\n", c.Name, foodName)
}

func main() {
	sleepList := []LazyAnimal{Dog{Name: "xiaobai"}, Cat{Name: "kitty"}}
	foodName := "food"
	for _, s := range sleepList {
		s.Sleep()
		s.Eat(foodName)
	}
}
```

大功告成，这里我们就实现了接口的嵌入，代码都是比较简单直白的。来总结一下重点内容：

- go 可以声明接口，它包含了一系列方法声明
- struct 可以实现接口，只要一个 struct 实现了一个接口的所有方法，我们就说 struct 实现了这个接口（隐式的）
- 接口也可以通过嵌入来声明一个新的接口，比如 ReadWriter 内嵌了 Reader 和 Writer。
- go 提倡“小而美”的接口，然后通过嵌入来组合新接口

## 类型断言(type assert)

上文我们看到在使用接口的地方，我们可以传入一个具体的实现了接口的 struct 类型，但是我们如何获取传入的到底是哪种 struct 类型呢？
go 给我们提供了一种方式叫做类型断言来获取具体的类型，它的语法比较简单，格式如下：

```go
instance, ok := interfaceVal.(RealType) // 如果 ok 为 true 的话，接口值就转成了我们需要的类型
```

我们继续再上边的代码里加上类型断言的演示，注意类型断言那几行代码，再 for 循环里边我们使用类型断言获取了接口值的真正类型。

```go
func main() {
	sleepList := []LazyAnimal{Dog{Name: "xiaobai"}, Cat{Name: "kitty"}}
	foodName := "food"
	for _, s := range sleepList {
		s.Sleep()
		s.Eat(foodName)

		// 类型断言 type assert
		if dog, ok := s.(Dog); ok {
			fmt.Printf("I am a Dog, my name is %s", dog.Name)
		}
		if cat, ok := s.(Cat); ok {
			fmt.Printf("I am a Cat, my name is %s", cat.Name)
		}
	}
}
```

## 使用空接口实现泛型

之前在函数那一章我们提到 go 目前没有直接提供对泛型的支持，学了接口之后其实我们可以用接口来实现。
上文提到，如果一个 struct 实现了一个接口声明所有方法，我们就说这个 struct (隐式)实现了这个接口，那如果是一个没有声明
任何方法的空接口(empty interface)呢？按照这个定义岂不是所有类型都实现了空接口么？

你猜对了，所有类型都实现了空接口(`interface{}`)，所以可以用空接口+类型断言转成任何我们需要的类型。来看下这个例子，
我们创建了一个空接口数组，它的元素可以是任何类型：

```go hl_lines="24"
package main

import (
	"fmt"
)

type Dog struct {
	Name string
}

func (d Dog) Sleep() {
	fmt.Printf("Dog %s is sleeping\n", d.Name)
}

type Cat struct {
	Name string
}

func (c Cat) Sleep() {
	fmt.Printf("Cat %s is sleeping\n", c.Name)
}

func main() {
	animalList := []interface{}{Dog{Name: "xiaobai"}, Cat{Name: "kitty"}}
	for _, s := range animalList {
		if dog, ok := s.(Dog); ok {
			fmt.Printf("I am a Dog, my name is %s\n", dog.Name)
		}
		if cat, ok := s.(Cat); ok {
			fmt.Printf("I am a Cat, my name is %s\n", cat.Name)
		}
	}
}
```

那我们如何实现泛型呢？空接口其实给了我们思路。既然它能转成所有类型，那我们以空接口作为参数不就好了嘛，这个想法是对的。
如果你有留意的话，到现在我们的代码示例里边使用最多的是啥，其实是这句话 `fmt.Println()`，不知道你之前有没有发现这个函数
居然可以传递任意类型进去，用的是什么黑魔法呢？

既然我们知道了空接口，可以自己实现一个简单的可以打印多种类型的 MyPrint 函数。

```go
func MyPrint(i interface{}) {
	switch o := i.(type) {
	case int:
		fmt.Printf("%d\n", o)
	case float64:
		fmt.Printf("%f\n", o)
	case string:
		fmt.Printf("%s\n", o)
	default:
		fmt.Printf("%+v\n", o)
	}
}

func main() {
	MyPrint(1)
	MyPrint(4.2)
	MyPrint("hello")
	MyPrint(map[string]string{"hello": "go"})
}
```

实际上如果你用开发工具跳转包 fmt 对应 `fmt.Println` 的函数实现，可以看到它也是以空接口作为参数的:

```go
// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Println(a ...interface{}) (n int, err error) {
	return Fprintln(os.Stdout, a...)
}
```

空接口在实现泛型的时候很有用，不过一般情况下如果不是必要，我们还是单独实现对应类型的函数就好，代码可读性也更高。

## 练习

- 实现一个简单的 `a+b` 函数，需要同时支持 int 和 float64 参数

## 参考

- [Golang and inheritance](https://stackoverflow.com/questions/32188653/golang-and-inheritance)
- [why generics](https://blog.golang.org/why-generics)
