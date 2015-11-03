package main

import (
	"fmt"
	"os"
)

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

type Stringer interface {
	String() string
}
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
func main() {

	var w Writer
	w = os.Stdout
	fmt.Fprintf(w, "Hello Word !\n")

	a := Person{"Deng", 24}
	b := Person{"Wang", 23}
	fmt.Println(a, b)

}
