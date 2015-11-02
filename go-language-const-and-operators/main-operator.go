package main

import (
	"fmt"
)

/*
6:  0110
11: 1011
---------
&	0010 = 2
|	1111 = 15
^	1101 = 13
&^	0100 = 4 if there is a 1 in right, the bit is 0, otherwise the bit is the same with left one
*/
func main() {
	fmt.Println(6 & 11)
	fmt.Println(6 | 11)
	fmt.Println(6 ^ 11)
	fmt.Println(6 &^ 11)
	fmt.Println(11 &^ 6)
}
