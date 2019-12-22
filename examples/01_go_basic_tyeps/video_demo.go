package main

import "fmt"

func testVar() {
	var i int64
	var s string
	fmt.Println("i is ", i)
	fmt.Println("s is ", s)

	ii := 88
	ss := "hello go"
	fmt.Println("ii is ", ii)
	fmt.Println("ss is ", ss)
}

func testBool() {
	var b1 bool
	b2 := true
	b3 := false
	fmt.Println(b1, b2, b3)
}

func testNumber() {
	var i int64
	var i2 int64
	fmt.Println(i, i2)
	i3 := 10
	i4 := int64(10)
	fmt.Println(int64(i3) + i4)

	// fmt.Println(1.0 / 0.0)
}

func testString() {
	s1 := "\"你好 Go\""
	s2 := `"你好 Go"`
	fmt.Println(s1)
	fmt.Println(s2)
}

const (
	Sunday = iota
	Monday
)

func main() {
	// testBool()
	// testNumber()
	// testString()
	fmt.Println(Sunday, Monday)
}
