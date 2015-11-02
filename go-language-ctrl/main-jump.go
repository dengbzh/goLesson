package main

import (
	"fmt"
)

//goto   work with tag, tag should be in upper case
//break
//continue
func main() {

	fmt.Println("start")

	goto JUMP

	fmt.Println("middle")

JUMP:
	fmt.Println("end")
LABLE:
	for {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			if i == 5 {
				break LABLE //use tag to break
			}
		}
	}
LABLE2:
	for a := 0; a < 9; a++ {
		for b := 0; b < 9; b++ {
			fmt.Println(a, b)
			if b == 5 {
				continue LABLE2 //use for outter loop
			}
		}
	}

	//continue cannot be used in switch, it only be used in loop
	/*
		switch a{
		case 0:
			fmt.Println("a=0");
			continue   //error
		}
	*/
	//continue match the loop
	a := 0
	for {
		switch a {
		case 0:
			fmt.Println("a=0")
			continue //match for
		}
		fmt.Println("continue") //not show
	}
}
