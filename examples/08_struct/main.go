package main

import "fmt"

type Animal struct {
	Name    string
	Age     int
	petName string
}

func (a Animal) Sleep() {
	fmt.Printf("%s is sleeping\n", a.Name)
}

func (a *Animal) SetPetName(petName string) {
	a.petName = petName // 注意这里的 a 是一个指针了
	// 以下这种方式也是可以的，go 如果碰到指针会自动帮我们处理，所以使用起来更方便
	// (*a).petName = petName
}

func main() {
	a := Animal{Name: "dog", Age: 3}
	a.SetPetName("hehe")
	a.Sleep()
	fmt.Println(a, a.Name, a.Age)

	// aPtr := &Animal{Name: "dog", Age: 3}
	// aPtr.SetPetName("little dog")
	// aPtr.Sleep()
	// fmt.Println(aPtr.petName) // 是不是可以设置成功了
}
