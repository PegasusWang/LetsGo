package main

import "fmt"

func testMap() {
	// 创建一个空的map
	m := make(map[string]int)
	m["hello"] = 1
	m["world"] = 2
	m["zhang"] = 3
	fmt.Println(m)
	// 输出一个存在的 key 和一个不存在的 key
	fmt.Println(m["hello"], m["not_found"])
	delete(m, "zhang") // 删除 zhang。注意删除一个不存在的key go 也不会抛出错误
	fmt.Println(m)
	if v, ok := m["hello"]; ok {
		fmt.Printf("m[%s] is %d \n", "hello", v)
	}
	// 同样使用 for/range 遍历，NOTE：遍历 map 返回的顺序是随机的，不要依赖 map 遍历的顺序
	for k, v := range m {
		fmt.Printf("m[%s]: %d\n", k, v)
	}
	// 如果只需要 k 或者 v 你可以使用 下划线作为占位符忽略值
	for k, _ := range m {
		fmt.Printf("k is %s\n", k)
	}
}

func testUseMapAsSet() {
	m := make(map[string]bool)
	m["hello"] = true
	m["world"] = true
	key := "hello"
	if _, ok := m[key]; ok {
		fmt.Printf("%s key exists\n", key)
	}
}

func main() {
	// testMap()
	testUseMapAsSet()
}
