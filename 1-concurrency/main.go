package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	numCount := 10
	sliceCh := make(chan int, numCount)
	powCh := make(chan int, numCount)

	go generateRandomNumbers(numCount, sliceCh)
	go calculateSquares(sliceCh, powCh)

	for v := range powCh {
		fmt.Println(v)
	}
}

func generateRandomNumbers(count int, out chan<- int) {
	for i := 0; i < count; i++ {
		out <- rand.Intn(101)
	}
	close(out)
}

func calculateSquares(in <-chan int, out chan<- int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}
