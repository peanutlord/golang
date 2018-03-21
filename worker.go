package main;

import (
	"fmt"
	"math/rand"
	"sync"
)

func emit()  <- chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 200; i++ {
			out <- rand.Intn(400);
		}
		close(out)
	}()
	return out
}

func worker(in <- chan int, wg *sync.WaitGroup, workerId int) {
	defer wg.Done()
	defer fmt.Printf("Worker %d done\n", workerId)

	fmt.Printf("Worker %d ready\n", workerId)
	for i := range in {
		fmt.Printf("Worker %d did some work: %d\n", workerId, i * i)
	}
}

func main() {
	emitter := emit()
	numWorker := 3

	var wg sync.WaitGroup
	wg.Add(numWorker)

	for i := 0; i < numWorker; i++ {
		go worker(emitter, &wg, i)
	}

	wg.Wait()
	fmt.Println("Done!")
}