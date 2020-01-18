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

# panic/recover


# 参考：

- [Go 面向失败编程](https://developer.aliyun.com/article/740696)
- [Error handling and Go](https://blog.golang.org/error-handling-and-go)
