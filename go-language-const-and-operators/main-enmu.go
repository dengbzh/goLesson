package main

import (
	"fmt"
)

const (
	a = "A"
	b // auto assignment b use expresstion of a
	c = "B"
	d
)

//iota conter
const (
	m = iota //default 0
	n        //auto increase
)

const (
	x = iota * 3 //works independetly with other iota
	y            //increase by 3
)

func main() {
	fmt.Println(a, b)
	fmt.Println(c, d)
	fmt.Println(m, n)
	fmt.Println(x, y)
}
