package main

import (
	"fmt"
	"time"
	"math/rand"
)

func emitter(prefix string) <- chan string {
	out := make(chan string)
	go func() {
		for i := 0; ; i++ {
			out <- fmt.Sprintf("Hello %s %d", prefix, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
		close(out)
	}()
	return out
}

func fan(emitter ...<-chan string) <- chan string {
	out := make(chan string)
	for _, e := range emitter {
		go func(single <- chan string) {
			for {
				out <- <-single
			}
		}(e)
	}

	return out
}

func main() {
	emitA := emitter("emitA")
	emitB := emitter("emitB")
	emitC := emitter("emitC")
	emitD := emitter("emitD")

	f := fan(emitA, emitB, emitC, emitD)
	for i := 0; i < 20; i++ {
		fmt.Println(<-f)
	}
}
