package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		fmt.Printf("g")
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	//fmt.Printf("past goroutine\n")
	// does NOT generate them all before printing the first, just often
	for range c {
		fmt.Printf("c")
	}
	fmt.Printf("\n") //done\n")
}
