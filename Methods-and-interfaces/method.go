/*
Methods can be associated with a named type or a pointer to a named type.

We just saw two Abs methods. One on the *Vertex pointer type and the other on the MyFloat value type.

*/

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type myFloat float64

func (f myFloat) Absf() float64 {
	if f < 0 {
		return float64(-f)
	} else {
		return float64(f)
	}
}
func main() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())

	f := myFloat(-math.Sqrt2)
	fmt.Println(f.Absf())
}
