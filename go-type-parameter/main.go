package main

import (
	"fmt"
)

//declaration
var a int

type Int bool

//multiple
var(
	aa int
	bb bool
	cc string = "hello"
	dd := 1.2
)

//with the same type
var n,m,k int = 1,2,3


func main() {
//without var 
	nn,mm,kk = 1.2,1,2
//sholdn't use var and ':' at the same time
	var b Int = true

	//auto
	//this cannot declar in global
	floatNum := 3.14

	// _=
	fmt.Println(a)
	a = 1
	fmt.Println(a)

	fmt.Println(b)
	fmt.Println(floatNum)
}
