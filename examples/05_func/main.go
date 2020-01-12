package main

import "fmt"

// 简单定义
func sum0(a int, b int) int {
	return a + b
}

// 相邻同类型可以省略前边参数的类型，只声明最后一个
func sum1(a, b int) int {
	return a + b
}

// 我们甚至还可以给返回值命名，这个时候需要通过赋值的方式来更新结果，而且 return 可以不用带返回值
func sum2(a, b int) (res int) {
	res = a + b
	return
}

// 可变参数
func sum3(init int, vals ...int) int {
	sum := init
	fmt.Println(vals, len(vals))
	for _, val := range vals {
		sum += val
	}
	return sum
}

// 返回多个值
func sum4(init int, vals ...int) (int, int) {
	sum := init
	fmt.Println(vals, len(vals))
	for _, val := range vals {
		sum += val
	}
	return sum, len(vals)
}

func testSum() {
	fmt.Println(sum0(1, 2))
	fmt.Println(sum1(1, 2))
	fmt.Println(sum2(1, 2))
	fmt.Println(sum3(0, 1, 2, 3))
	fmt.Println(sum4(0, 1, 2, 3))
}
func changeStr(s string) {
	s = "hehe"
	fmt.Println(s)
}

func test_changeStr() {
	name := "lao wang"
	changeStr(name)
	fmt.Println(name) // 打印出来还是 "lao wang"，没有修改成功
}
func changeMap(m map[string]string) {
	m["王八"] = "绿豆"
}

func changeString(s *string) {
	*s = "new lao wang"
}

func main() {
	s := "lao wang"
	changeString(&s)
	fmt.Println(s)
	// m := map[string]string{"name": "lao wang"}
	// changeMap(m)
	// fmt.Println(m) // map[name:lao wang 王八:绿豆]
}
