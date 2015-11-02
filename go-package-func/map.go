package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

func main() {

	m := make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.22, 13.03,
	}

	fmt.Println(m["Bell Labs"])

	n := make(map[string]int)

	n["Deng"] = 100
	fmt.Println("The value ", n["Deng"])

	n["Deng"]++
	fmt.Println("The value ", n["Deng"])

	delete(n, "Deng")
	fmt.Println("The value ", n["Deng"])

	v, ok := n["Deng"]
	fmt.Println("The value ", v, "Present?", ok)

	var k = map[string]Vertex{
		"Lab": Vertex{
			2.1, 2.2,
		},
		"Google": Vertex{
			3.1, 3.2,
		},
	}
	fmt.Println(k)

	var x = map[string]Vertex{
		"Bai": {1.1, 2.1},
		"Du":  {1.0, 0.1},
	}
	fmt.Println(x)
}
