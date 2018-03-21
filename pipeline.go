package main

import (
	"fmt"
	"time"
)

func gen() <- chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Printf("Emmited %d\n", i)
			out <- i
		}
		close(out)
	}();

	return out
}

func multiply(in <- chan int) <- chan int {
	out := make(chan int)
	go func() {
		for i := range in {
			fmt.Printf("Calculated %d from %d\n", i*i, i)
			out <- i*i
		}
	}()

	return out
}

func print(in <- chan int) {
	go func() {
		for _ = range in {
			fmt.Println(<-in)
		}
	}()
}

func main() {
	num := gen()
	mul := multiply(num)
	print(mul)

	time.Sleep(10 * time.Second)
}