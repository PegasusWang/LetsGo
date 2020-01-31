package main

import "fmt"

// 自定义一个 Enum类型
type Enum int

const (
	// 这里如果是自增用 iota 更好
	Init    Enum = 0
	Success Enum = 1
	Fail    Enum = 2

	// 枚举对应的中文
	InitName    = "初始化"
	SuccessName = "成功"
	FailName    = "失败"
)

func (e Enum) Int() int {
	return int(e)
}

func (e Enum) String() string {
	return []string{
		InitName,
		SuccessName,
		FailName,
	}[e]
}

func testEnum() {
	status := 0
	fmt.Println(Init.Int() == status)

	status2 := Fail
	fmt.Println(status2.String())
}

func main() {
	// 定义一个 counter 类型
	type Counter map[string]int
	c := Counter{}
	c["word"]++
	fmt.Println(c)

	type Queue []int
	q := make(Queue, 0)
	q = append(q, 1)
	fmt.Println(q)
}
