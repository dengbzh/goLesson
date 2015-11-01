package main

//first line is package unless it's a comment
import (
	"mft"
)

const (
	PI   float64 = 3.14
	NUM  int64   = 2
	NAME string  = "deng"
)

var (
	Pi   float64 = 3.14
	Num  int64   = 2
	Name string  = "deng"
)

type Gopher struct {
}

func (gopher Gopher) Go() {

}

type Writer interface {
}

//it's my main function
/*

 */
func main() {
	Println("haha")

}

//comment
/*
	comment

*/
