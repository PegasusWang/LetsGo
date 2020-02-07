# go 是如何处理错误的？

!!! quote
    "You can't write perfect software. Software can't be perfect. Protect your code and users from the inevitable errors.

## 错误处理

如果你使用过 Python/Java 等流行的编程语言，你会发现它们使用异常机制来进行错误处理，Python 中你可以使用 try/except
来进行异常的捕获和处理，如果异常没有被捕获，就会造成程序退出。还可以通过异常栈来追踪异常的调用信息从而帮助我们修复异常代码。

在 go 中使用的是类似 c 的返回错误的方式，比如我们在 go 代码中经常会看到很多这种错误检查代码。go
的这种比较原始的错误处理方式实际上也招来了很多人的批评，认为是设计上的一个败笔，很多习惯使用了异常机制的开发者可能会
感觉 go 的错误处理写起来非常繁琐，比如下面这个例子，每个流程你都要去处理返回的错误值。

```go
// from https://8thlight.com/blog/kyle-krull/2018/08/13/exploring-error-handling-patterns-in-go.html
func (router HttpRouter) parse(reader *bufio.Reader) (Request, Response) {
  requestText, err := readCRLFLine(reader) //string, err Response
  if err != nil {
    //No input, or it doesn't end in CRLF
    return nil, err
  }

  requestLine, err := parseRequestLine(requestText) //RequestLine, err Response
  if err != nil {
    //Not a well-formed HTTP request line with {method, target, version}
    return nil, err
  }

  if request := router.routeRequest(requestLine); request != nil {
    //Well-formed, executable Request to a known route
    return request, nil
  }

  //Valid request, but no route to handle it
  return nil, requestLine.NotImplemented()
}
```

在 go 的惯例中，一般函数多个返回值的最后一个值用来返回错误，返回 nil 表示没有错误，调用者通过检查返回的错误是否是 nil
就知道是否需要处理错误了。

## defer 语句

go 中提供了一个 defer 语句用来延迟一个函数(匿名函数)或者方法的执行，它会在函数执行完成之后调用。一般为了防止代码里有资源泄露，
对于打开的资源比如文件等我们需要显示进行关闭，这种场合就是 defer 发挥作用最好的场景，也是 go 代码中使用 defer 最常用的场景。

``` go tab="Go"
f, err := os.Open(file)
if err != nil {
  // handle err
  return err
}
defer f.Close() // 保证文件会在函数返回之后关闭，防止资源泄露
```

``` py tab="Python"
with open("filepath", "r") as f:
    # do with file
```

如果你用过 python 的话，go 中的 defer 和 python 使用 with 语句保证资源会被关闭目的一样。
另外函数里可以使用多个 defer 语句，如果有多个 defer 它们会按照后进先出(Last In First Out)的顺序执行。
运行以下小例子，看看输出是否和你想的一样：

```go
package main

import (
	"fmt"
)

func testDefer() string {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	fmt.Println("函数体")
	return "hello"
}

func main() {
	fmt.Println(testDefer())
}
```

## go 的 error 类型

上文提到一般我们在 go 中通过返回一个 error 来表示错误或者异常状态，这是 go 代码中最常见的方式。那 error 究竟是什么呢？
其实 error 是 go 的一个内置的接口类型，比如你可以使用开发工具跳进去看下 error 的定义（注意这里使用到了接口，后面会介绍）。

```go
// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
type error interface {
	Error() string
}
```

error 的定义很简单，只要我们自己实现了一个类型的 Error() 方法返回一个字符串，就可以当做错误类型了。举一个简单小例子，
比如计算两个整数相除，我们知道除数是不能为 0 的，这个时候我们就可以写个函数：

```go
import (
	"errors" // 使用内置的 errors
	"fmt"
)

// Divide compute int a/b
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}

func main() {
	// fmt.Println(testDefer())
	a, b := 1, 0
	res, err := Divide(a, b)
	if err != nil {
		fmt.Println(err) // error 类型实现了 Error() 方法可以打印出来
	}
	fmt.Println(res)
}
```

在内置库和业务代码中，你会看到很多类似的处理方式，比如我跳进去到 fmt 包里随便搜到的一个例子(看内置 package 可以学到很多惯用法)：

```go
func (p *pp) Write(b []byte) (ret int, err error) {
	p.buf.Write(b)
	return len(b), nil
}
```

在我们的业务代码中也是这样，如果你希望返回一个错误，可以在函数的最后一个返回值返回一个错误类型。

## 自定义自己的业务异常 TODO

## Go 的异常处理 panic/recover

上边我们提到了错误，这里聊聊 go 的异常处理机制 panic(恐慌)/recover(恢复)，其实一般我们使用的是错误处理(error)而不是 panic。因为只有非常严重的场景
下才会发生 panic 导致代码退出。平常我们使用的 web 框架，一般即使出错了，我们也希望整个进程继续执行，而不是直接退出无法处理用户请求。
比如 python 的 web 框架，如果遇到了业务代码没有捕获的异常，框架会给我们捕获然后返回给客户端 500 的状态码表示代码有错。

go 里区分对待异常(panic)和错误(error)的，绝大部分场景下我们使用的都是错误，只有少数场景下发生了严重错误我们想让整个进程都退出了才会使用异常。

举个例子，在 web 框架启动之前经常需要读取配置文件，获取 mysql/redis 等服务的地址和端口，这个时候如果读取配置的代码失败了，
我会使用 panic 直接退出，这就是严重错误，因为即使服务启动了也无法连接到数据库正常处理请求，反而直接退出进程早发现错误为好。
但是比如用户传了一个错误参数这种情况不严重的错误，我们希望直接给用户返回一个错误状态码，而不至于退出进程。

```go
if err := readConfig("filepath"); err != nil {
	panic(err) // 读取失败直接退出
}
```

比如刚才除法函数的例子，如果我们碰到了个除数为 0 被认为是严重错误，也可以使用 panic 抛出异常：

```go
func MustDivide(a, b int) int {
	if b == 0 {
		panic("divide by zero")
	}
	return a / b
}
```

如果我们不幸传入了除数为0，但是又不想让进程退出呢？go 还提供了一个 recover 函数用来从异常中恢复，比如使用 recover
可以把一个 panic 包装成为 error 再返回，而不是让进程退出：

```go
func Divide2(a, b int) (res int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	res = MustDivide(a, b)
	return res, nil
}
```

这样一来我们就『捕获』了 panic 异常并且返回了一个错误，代码也可以正常执行而不会退出啦。

最后简单总结一下：

- 对于一般不太严重的场景，返回错误值 error 类型 (业务绝大部分场景)
- 对于严重的错误需要整个进程退出的场景，使用 panic 来抛异常，及早发现错误
- 如果希望捕获 panic 异常，可以使用 recover 函数捕获，并且包装成一个错误返回
- web 框架等会帮你捕获 panic 异常，然后返回客户端一个 http 500 状态码错误

文末还有几篇不错的文章供大家参考，多看一些流行的源码，你就知道它们的使用场合了。

## 参考：

- [Go 面向失败编程](https://developer.aliyun.com/article/740696)
- [Error handling and Go](https://blog.golang.org/error-handling-and-go)
- [Defer, Panic, and Recover](https://blog.golang.org/defer-panic-and-recover)
