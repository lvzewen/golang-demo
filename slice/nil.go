package main

import (
	"fmt"
)

// C ...
type C struct {
	Tsst string
}

func main()  {
	var a []int
	var b []*int
	var c []*C

	fmt.Println("a == nil", a == nil)
	fmt.Println("b == nil", b == nil)
	fmt.Println("c == nil", c == nil)
}