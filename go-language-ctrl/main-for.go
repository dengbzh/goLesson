package main

import (
	"fmt"
)

func num() int {
	fmt.Println("number")
	return 3
}
func main() {
	a := 1
	//infinite loop
	for {
		a++
		if a != 5 {
			fmt.Println(a)
		} else {
			break //leave loop
		}
	}
	a = 1
	//condition
	for a <= 5 {
		fmt.Println(a)
		a++
	}
	//inital and step forward
	for a = 1; a <= 5; a++ {
		fmt.Println(a)
	}
	//multiple
	for a, b := 1, 2; a+b < 5; a, b = a+1, b+1 { //cannot use a++ and b++
		fmt.Println(a, b)
	}
	//use var to reduce function invoke
	for a := 1; a < num(); a++ {
		fmt.Println(a)
	}
	var n = num() //invoke num() once
	for a := 1; a < n; a++ {
		fmt.Println(a)
	}

}
