package main

import (
	"fmt"
)

/*
By default, sends and receives block until the other side is ready.

This allows goroutines to synchronize without explicit locks or condition variables.
*/
//channel is a fifo
func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 6}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}
