package main

import (
	"fmt"
)

type Dog struct {
	Name string
}

func (d Dog) Sleep() {
	fmt.Printf("Dog %s is sleeping\n", d.Name)
}

type Cat struct {
	Name string
}

func (c Cat) Sleep() {
	fmt.Printf("Cat %s is sleeping\n", c.Name)
}

func testEmptyInterface() {
	animalList := []interface{}{Dog{Name: "xiaobai"}, Cat{Name: "kitty"}}
	for _, s := range animalList {
		if dog, ok := s.(Dog); ok {
			fmt.Printf("I am a Dog, my name is %s\n", dog.Name)
		}
		if cat, ok := s.(Cat); ok {
			fmt.Printf("I am a Cat, my name is %s\n", cat.Name)
		}
	}
}

func MyPrint(i interface{}) {
	switch o := i.(type) {
	case int:
		fmt.Printf("%d\n", o)
	case float64:
		fmt.Printf("%f\n", o)
	case string:
		fmt.Printf("%s\n", o)
	default:
		fmt.Printf("%+v\n", o)
	}
}

func main() {
	MyPrint(1)
	MyPrint(4.2)
	MyPrint("hello")
	MyPrint(map[string]string{"hello": "go"})
}
