package main

import "fmt"

func testArray() {
	var arrayInt64 [3]int64
	arrayInt64[0], arrayInt64[1], arrayInt64[2] = 0, 1, 2
	fmt.Println(arrayInt64)

	arrayString := []string{"hello", "go"}
	fmt.Println(arrayString)

	arrayFloat := [...]float64{1.0, 2.3, 3.3}
	fmt.Println(arrayFloat)

	matrix := [2][2]int64{
		{0, 1},
		{2, 3},
	}
	fmt.Println(matrix)
}
func testSlice() {
	ints := make([]int, 3) // 默认零值
	ints[0], ints[1] = 1, 2
	fmt.Println(ints)

	names := []string{"zhang", "wang", "li", "zhao"}
	fmt.Println(names, len(names), cap(names))
	names2 := names[0:3]
	fmt.Println(names2)
	names2[0] = "lao zhang"
	fmt.Println(names)

	for _, name := range names {
		fmt.Println(name)
	}

	vals := make([]int, 0)
	for i := 0; i < 3; i++ {
		vals = append(vals, i)
	}
	fmt.Println(vals)
	vals2 := []int{3, 4, 5}
	newVals := append(vals, vals2...)
	fmt.Println(newVals)
}

func main() {
	// testArray()
	testSlice()
}
