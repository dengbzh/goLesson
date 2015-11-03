package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("invalid input negative numbers")
}

func Sqrt(x ErrNegativeSqrt) (float64, error) {
	if x < 0 {
		return -1, x
	}
	return 0, nil
}

func main() {
	fmt.Println((Sqrt(2)))
	fmt.Println((Sqrt(-2)))
}
