/*
Interfaces
An interface type is defined by a set of methods.

A value of interface type can hold any value that implements those methods.

Note: There is an error in the example code on line 22. Vertex (the value type) doesn't satisfy Abser because	the Abs method is defined only on *Vertex (the pointer type).
*/

package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {

	var a Abser
	f := myFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	a = &v
	//a = v

	fmt.Println(a.Abs())
}

type myFloat float64

func (f myFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
