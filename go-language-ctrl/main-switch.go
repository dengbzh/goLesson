package main

import (
	"fmt"
)

func main() {
	a := 0
	switch a { //
	case 0:
		fmt.Println(0)
		fallthrough //to run next one case
	case 1:
		fmt.Println(1)
		fallthrough
	case 2:
		fmt.Println(2)
		break
		fmt.Println("break")
	case 3, 4:
		fmt.Println("3,4")
	default:
		fmt.Println("default")
	}
	//without condition
	switch {
	case a >= 0:
		fmt.Println("a>=0")
	case a >= 2:
		fmt.Println("a>=2")
	default:
		fmt.Println("default")
	}
}
