package main

import (
	"fmt"
)

func main() {
	intChain := make(chan int)
	go func ()  {
		intChain <- 1
	} ()
	value := <- intChain
	fmt.Println("value: ", value)
}