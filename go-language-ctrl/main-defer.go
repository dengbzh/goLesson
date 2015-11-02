package main

import (
	"fmt"
)

func main() {

	defer fmt.Println("world") //this will execute after next func return

	fmt.Println("hello")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
