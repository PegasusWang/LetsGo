package main

import (
	"fmt"
	"math"
)

func boolDemo() {
	var b1 bool
	b2 := false
	b3 := true
	if b1 {
		fmt.Println("b1 is true")
	}
	if b2 {
		fmt.Println("b2 is true")
	}
	if b3 {
		fmt.Println("b3 is true")
	}
}

func intDemo() {
	var i64 int64
	i64 = 10
	fmt.Println(i64 + 10)

	i32 := int32(42)
	fmt.Println(i32 + 10)

	fmt.Println(
		math.MaxInt64,
	)

	a, b := 10, 0
	fmt.Println(a / b) // Boom! 会发生什么

}

func stringDemo() {
	// 如果字符串本身也有双引号，就需要把里边的使用 \ 转义
	s1 := "\"Hello Go\""
	// 使用反斜线就可以直接包含双引号了
	s2 := `"Hello Go"`
	fmt.Println(s1) // 打印出 "Hello Go"
	fmt.Println(s2) // 打印出 "Hello Go"

	s3 := `
你好
`
	s4 := "Golang !"
	fmt.Println(s3 + s4)

	hs := "你好 Golang"
	for idx, c := range hs {
		fmt.Println(idx, string(c))
	}

}

func identifierDemo() {
	var i int64             // 声明一个 int64 变量
	var b string            // 声明一个字符串
	fmt.Println("i is ", i) // 0
	fmt.Println("b is ", b) // ""

	// 还有一种简化方式，声明并且赋值，编译器负责推断类型
	ii := 1
	s := "Hello Go!"
	fmt.Println("ii is ", ii) // 1
	fmt.Println("s is ", s)   // Hello Go!"
}

func main() {
	// boolDemo()
	// intDemo()
	stringDemo()
}
