package main

import "fmt"

func testIf() {
	ok := true
	if ok {
		fmt.Println("ok is true")
	}

	day := "Friday"
	if day == "Friday" {
		fmt.Println("明天不上班呀!")
	} else if day == "Sunday" {
		fmt.Println("周末好快")
	} else {
		fmt.Println("干活啦")
	}

	m := make(map[string]string)
	m["王八"] = "绿豆"
	if v, ok := m["王八"]; ok {
		fmt.Println(v)
	}

}
func testSwitch() {
	// 常规用法
	day := 0
	switch day {
	case 0, 7:
		fmt.Println("周末")
	case 1, 2, 3, 4, 5:
		fmt.Println("工作日")
	default:
		fmt.Println("不合法")
	}
	// case 后边还可以是表达式
	a, b := 1, 2
	a, b = b, a
	switch {
	case a < b:
		fmt.Println("a < b")
	case a > b:
		fmt.Println("a > b")
	}

}

func testFor() {
	intSlice := []int{3, 2, 1}
	for index, item := range intSlice {
		fmt.Println(index, item)
	}
	for index := range intSlice { // 省略 item 之后遍历的是 key，注意不像python 直接遍历值
		fmt.Println(index)
	}

	m := map[string]string{"k1": "v1", "k2": "v2"}
	for k, v := range m {
		fmt.Println(k, v)
	}
	for k := range m {
		fmt.Println(k)
	}
}

func main() {
	// testIf()
	// testSwitch()
	testFor()
}
