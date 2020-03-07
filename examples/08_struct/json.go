package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Animal struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	petName string `json:"pet_name"`
}

func main() {
	animals := []Animal{
		Animal{"dog", 3, "little dog"},
		Animal{"cat", 4, "little cat"},
	}
	bs, err := json.Marshal(animals)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bs)) // [{"name":"dog","age":3},{"name":"cat","age":4}]
}
