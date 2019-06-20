package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	b := &B{
		Test: "test",
		Test1: "test1",
	}

	tmp, _ := json.Marshal(b)

	a := &A{}

	_ = json.Unmarshal(tmp, a)

	fmt.Println("a", a)
}

// A ...
type A struct {
	Test string `json:"test"`
}

// B ...
type B struct {
	Test string `json:"test"`
	Test1 string `json:"test1"`
}
