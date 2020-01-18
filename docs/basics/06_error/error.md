> "You can't write perfect software. Software can't be perfect. Protect your code and users from the inevitable errors.

# 错误处理

如果你使用过 Python/Java 等流行的编程语言，你会发现它们使用异常机制来进行错误处理，Python 中你可以使用 try/except
来进行异常的捕获和处理，如果异常没有被捕获，就会造成程序退出。还可以通过异常栈来追踪异常的调用信息从而帮助我们修复异常代码。

在 go 中使用的是类似 c 的返回错误值的方式，比如我们在 go 代码中经常会看到很多这种错误检查代码。go
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

在 go 的惯例中，一般多个返回值的最后一个值用来返回错误，返回 nil 表示没有错误，调用者通过检查返回的错误是否是 nil
就知道是否需要处理错误了。

# defer 语句

go 中提供了一个 defer 语句用来延迟一个函数(匿名函数)或者方法的执行，它会在函数执行完成之后调用。一般为了防止代码里有资源泄露，
对于打开的资源比如文件等我们需要显示就行关闭，这种场合就是 defer 发挥作用最好的场景，也是 go 代码中使用 defer 最常用的场景。

```go
f, err := os.Open(file)
if err != nil {
  // handle err
  return err
}
defer f.Close() // 保证文件会在函数返回之后关闭，防止资源泄露
```

其实如果你用过 python 的话，和 python 使用 with 语句保证资源会被关闭目的一样。
另外函数里可以使用多个 defer 语句，如果有多个 defer 它们会按照后进先出(Last In First
Out)的顺序执行。运行以下小例子，看看输出是否和你想的一样：

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

# go 的 error 类型

上文提到一般我们在 go 中通过返回一个 error 来表示错误或者异常状态，这是 go 代码中最常见的方式。那 error 究竟是什么呢？
其实 error 是 go 的一个内置的接口类型，比如你可以使用开发工具跳进去看下 error 的定义（注意这里使用到了接口，后面会介绍）。

```go
// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
type error interface {
	Error() string
}
```

error 的定义很简单，只要我们自己实现了一个类型的 Error() 方法返回一个字符串，就可以当做错误类型了。举一个简单小例子，比
如计算两个整数相除，我们知道除数是不能为 0 的，这个时候我们就可以写个函数：

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

# Go 的异常处理 panic/recover

上边我们提到了错误，这里聊聊 go 的异常处理机制 panic(恐慌)/recover(恢复)，其实一般我们使用的是错误处理(error)而不是 panic。因为只有非常严重的场景
下才会发生 panic 导致代码退出。go 里是区分对待异常(panic)和错误(error)类型的


# 参考：

- [Go 面向失败编程](https://developer.aliyun.com/article/740696)
- [Error handling and Go](https://blog.golang.org/error-handling-and-go)
- [Defer, Panic, and Recover](https://blog.golang.org/defer-panic-and-recover)
