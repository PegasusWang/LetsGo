package main

import "fmt"

func testArray() {
	// 声明一个 int64数组，声明之后每个元素是该类型默认的『零值』
	var arrayIn64 [3]int64
	arrayIn64[0], arrayIn64[1] = 1, 2
	fmt.Println(arrayIn64)
	// 声明并且初始化
	arrayString := [3]string{"zhang", "wang", "li"}
	fmt.Println(arrayString)
	// 也可以省略长度，让 go 自动计算。这个时候你需要是使用省略号 ...
	// 创建一个长度为 3 的 float64 数组
	arrayFloat := [...]float64{1.5, 8.8, 6.6}
	fmt.Println(arrayFloat)
	// 二维数组
	matrix := [2][2]int64{
		{0, 1},
		{2, 3},
	}
	fmt.Println(matrix)
}

func testArrayOperation() {
	names := [4]string{"zhang", "wang", "li", "zhao"}
	fmt.Printf("names has %d elements\n", len(names))
	fmt.Println(names[1]) // NOTE: 注意如果下标超过范围会 panic
	names[3] = "lao zhang"
	fmt.Println(names[3]) // 赋值
	// 获取区间
	fmt.Println(names[2:4])
}

func testSlice() {
	// 创建并且初始化一个 slice
	names := []string{"zhang", "wang", "li", "zhao"}
	// 打印 names, 长度和容量
	fmt.Println(names, len(names), cap(names))
	names2 := names[0:3] // 获取子切片 0,1,2 三个元素，注意左闭右开区间
	fmt.Println(names2)
	// 尝试修改一下 names2 ，注意 names 也会跟着改变么？
	names[0] = "lao zhao"
	fmt.Println(names, names2) // 你会发现names也变了，这里起始它们共用了底层结构，注意这个问题

	// 遍历一个 slice 我们使用 for/range 语法
	for idx, name := range names { // 如果没有用到下标 idx，可以写成下划线 _ 作为占位符，但是不能省略
		fmt.Println(idx, name)
	}

	// 修改切片主要通过赋值和 append 操作。使用 append 修改切片
	vals := make([]int, 0)
	for i := 0; i < 3; i++ {
		vals = append(vals, i)
	}
	fmt.Println(vals)
	vals2 := []int{3, 4, 5}
	newVals := append(vals, vals2...) // 可以使用省略号的方式『解包』一个 slice 来连接两个 slice
	fmt.Println(newVals)
}

func main() {
	// testArray()
	// testArrayOperation()
	testSlice()
}
