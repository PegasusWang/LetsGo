package main

import "fmt"

func testMap() {
	m := make(map[string]int)
	m["hello"] = 1
	m["world"] = 2
	m["zhang"] = 3
	fmt.Println(m)
	fmt.Println(m["hello"], m["not_found"])
	key := "hello"
	if v, ok := m[key]; ok {
		fmt.Println("key is exists", key, v)
	}
	delete(m, "zhang")

	m["wang"] = 8
	m["laoli"] = 9
	for k, v := range m {
		fmt.Printf("m[%s] : %d\n", k, v)
	}
	for key := range m {
		fmt.Println(key)
	}
	for _, v := range m {
		fmt.Println(v)
	}

}

func main() {
	testMap()
}
