package main

import (
	"fmt"
)

func main() {
	s := []int{2, 3, 4, 5, 6, 7}
	fmt.Println("s ==", s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}

	fmt.Println("s[1:4] ==", s[1:4])
	//from s[0]
	fmt.Println("s[:3] ==", s[:3])
	//until s[(len)]
	fmt.Println("s[4:] ==", s[4:])
}
