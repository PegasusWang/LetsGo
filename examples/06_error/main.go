package main

import (
	"errors"
	"fmt"
)

func testDefer() string {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	fmt.Println("函数体")
	return "hello"
}

// Divide compute int a/b
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}

func MustDivide(a, b int) int {
	if b == 0 {
		panic("divide by zero")
	}
	return a / b
}

func Divide2(a, b int) (res int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	res = MustDivide(a, b)
	return res, nil
}

func testDivide() {
	// fmt.Println(testDefer())
	a, b := 1, 0
	res, err := Divide(a, b)
	if err != nil {
		fmt.Println(err) // error 类型实现了 Error() 方法可以打印出来
	}
	fmt.Println(res)
}

func main() {
	// MustDivide(1, 0)
	fmt.Println(Divide2(1, 0))
}
