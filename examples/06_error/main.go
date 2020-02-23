package main

import "fmt"

func testDefer() string {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	fmt.Println("函数体")
	return "hello"
}

type ArticleError struct {
	Code    int32
	Message string
}

func (e *ArticleError) Error() string {
	return fmt.Sprintf("[ArticleError] Code=%d, Message=%s", e.Code, e.Message)
}

func NewArticleError(code int32, message string) error {
	return &ArticleError{
		Code:    code,
		Message: message,
	}
}

func MustDivide(a, b int) int {
	if b == 0 {
		panic("divide by zero")
	}
	return a / b
}

func Divide2(a, b int) (res int, e error) {
	defer func() {
		if err := recover(); err != nil { // panic -> error
			e = fmt.Errorf("%v", err)
		}
	}()
	res = MustDivide(a, b)
	return
}

func main() {
	// fmt.Println(testDefer())
	// MustDivide(1, 0)
	// fmt.Println("end")
	res, err := Divide2(10, 0)
	fmt.Println(res, err)
}
