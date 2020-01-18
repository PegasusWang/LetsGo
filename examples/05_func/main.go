package main

import (
	"fmt"
)

func sum0(a int, b int) int {
	return a + b
}

func sum1(a, b int) int {
	return a + b
}

func sum2(a, b int) (res int) {
	res = a + b
	return
}

func sum3(init int, vals ...int) int {
	sum := init
	for _, val := range vals {
		sum += val
	}
	return sum
}

func sum4(init int, vals ...int) (int, int) {
	sum := init
	for _, val := range vals {
		sum += val
	}
	return sum, len(vals)
}

func main() {
	fmt.Println(sum0(1, 2))
	fmt.Println(sum1(1, 2))
	fmt.Println(sum2(1, 2))
	fmt.Println(sum3(0, 1, 2, 3))
	fmt.Println(sum3(0, []int{1, 2, 3}...))
	fmt.Println(sum4(0, []int{1, 2, 3}...))
	sum, length := sum4(0, []int{1, 2, 3}...)
	fmt.Println(sum, length)
}
