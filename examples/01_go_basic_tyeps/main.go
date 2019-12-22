package main

import (
	"fmt"
	"math"
	"strconv"
)

const (
	Sunday    = 0
	Monday    = 1
	Tuesday   = 2
	Wednesday = 3
	Thursday  = 4
	Friday    = 5
	Saturday  = 6
)

func testIota() {
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Monday)
}

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

func testConvert() { // 测试 int 和 string(decimal) 互相转换的函数
	// https://yourbasic.org/golang/convert-int-to-string/
	// int -> string
	sint := strconv.Itoa(97)
	fmt.Println(sint, sint == "97")

	// byte -> string
	bytea := byte(1)
	bint := strconv.Itoa(int(bytea))
	fmt.Println(bint)

	// int64 -> string
	sint64 := strconv.FormatInt(int64(97), 10)
	fmt.Println(sint64, sint64 == "97")

	// int64 -> string (hex) ，十六进制
	sint64hex := strconv.FormatInt(int64(97), 16)
	fmt.Println(sint64hex, sint64hex == "61")

	// string -> int
	_int, _ := strconv.Atoi("97")
	fmt.Println(_int, _int == int(97))

	// string -> int64
	_int64, _ := strconv.ParseInt("97", 10, 64)
	fmt.Println(_int64, _int64 == int64(97))

	// https://stackoverflow.com/questions/30299649/parse-string-to-specific-type-of-int-int8-int16-int32-int64
	// string -> int32，注意 parseInt 始终返回的是 int64，所以还是需要 int32(n) 强转一下
	_int32, _ := strconv.ParseInt("97", 10, 32)
	fmt.Println(_int32, int32(_int32) == int32(97))

	// int32 -> string, https://stackoverflow.com/questions/39442167/convert-int32-to-string-in-golang
	i := 42
	strconv.FormatInt(int64(i), 10) // fast
	strconv.Itoa(int(i))            // fast
	fmt.Sprint(i)                   // slow

	// int -> int64 ，不会丢失精度
	var n int = 97
	fmt.Println(int64(n) == int64(97))

	// string -> float32/float64  https://yourbasic.org/golang/convert-string-to-float/
	f := "3.14159265"
	if s, err := strconv.ParseFloat(f, 32); err == nil {
		fmt.Println(s) // 3.1415927410125732
	}
	if s, err := strconv.ParseFloat(f, 64); err == nil {
		fmt.Println(s) // 3.14159265
	}
}

// func main() {
// 	// boolDemo()
// 	// intDemo()
// 	// stringDemo()
// 	testIota()
// }
