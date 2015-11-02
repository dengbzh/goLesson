/*
if
*/
package main

import (
	"fmt"
)

func main() {
	a := 1
	b := 2
	/*
		if a = 3; a == 3 { //initial a = 3
			fmt.Println("a=3")
		}*/
	if a := 3; a == 3 { //creat local var a
		fmt.Println("a is local varible", a)
	}
	fmt.Println("it's not the same a", a)

	if a, b := 1, 2; a < b { //initial multiple var
		fmt.Println(a, b)
	}
	if a < b { //this brace must in the same line with the last condition
		fmt.Println("a<b")
	} else if a > b &&
		a == b {
		fmt.Println("a>b")
	}
}
