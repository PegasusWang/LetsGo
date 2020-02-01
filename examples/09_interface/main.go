package main

import (
	"fmt"
)

// Sleeper 接口声明
type Sleeper interface {
	Sleep() // 声明一个 Sleep() 方法
}

type Eater interface {
	Eat(foodName string)
}
type LazyAnimal interface {
	Sleeper
	Eater
}

type Dog struct {
	Name string
}

func (d Dog) Sleep() {
	fmt.Printf("Dog %s is sleeping\n", d.Name)
}

func (d Dog) Eat(foodName string) {
	fmt.Printf("Dog %s is eating %s\n", d.Name, foodName)
}

type Cat struct {
	Name string
}

func (c Cat) Sleep() {
	fmt.Printf("Cat %s is sleeping\n", c.Name)
}

func (c Cat) Eat(foodName string) {
	fmt.Printf("Cat %s is eating %s\n", c.Name, foodName)
}

func main() {
	sleepList := []LazyAnimal{Dog{Name: "xiaobai"}, Cat{Name: "kitty"}}
	foodName := "food"
	for _, s := range sleepList {
		s.Sleep()
		s.Eat(foodName)
	}
}
