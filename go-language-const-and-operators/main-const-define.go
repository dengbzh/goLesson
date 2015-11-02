package main

import (
	"fmt"
)

const a int = 1
const b = 2

const (
	text   = "123"
	length = len(text)
	num    = a * 20
)

const i, j, k = 10, "20", 30

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(text, length, num)
}
