package main

import "fmt"

type Enum int

const (
	Init    Enum = 0
	Success Enum = 1
	Fail    Enum = 2

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

func main() {
	// fmt.Println(Success)
	type Counter map[string]int
	c := Counter{}
	c["word"]++
	fmt.Println(c)

	type Queue []int
	q := make(Queue, 0)
	q = append(q, 1)
	fmt.Println(q)
}
