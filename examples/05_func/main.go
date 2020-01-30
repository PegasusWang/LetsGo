package main

import (
	"fmt"
)

func changeStr(s string) {
	s = "hehe"
	fmt.Println(s)
}
func changeStrByPtr(s *string) {
	*s = "new hehe"
}

func changeMap(m map[string]string) {
	m["王八"] = "绿豆"
}

func FilterIntSlice(intVals []int, predicate func(i int) bool) []int {
	res := make([]int, 0)
	for _, val := range intVals {
		if predicate(val) {
			res = append(res, val)
		}
	}
	return res
}

func main() {
	name := "lao wang"
	changeStr(name) // 值传递
	fmt.Println(name)

	m := map[string]string{"name": "wang"}
	changeMap(m)
	fmt.Println(m)

	name2 := "name2"
	changeStrByPtr(&name2)
	fmt.Println(name2)

	ints := []int{1, 2, 3, 4, 5}
	isEven := func(i int) bool { return i%2 == 0 }
	fmt.Println(FilterIntSlice(ints, isEven))
}
