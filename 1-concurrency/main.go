package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.NewSource(time.Now().UnixNano())

	numCount := 10
	sliceCh := make(chan int, numCount)
	powCh := make(chan int, numCount)

	go func() {
		slice := make([]int, numCount)
		for i := 0; i < numCount; i++ {
			slice[i] = rand.Intn(101)
			sliceCh <- slice[i]
		}
		close(sliceCh)
	}()

	go func() {
		for v := range sliceCh {
			powCh <- v * v
		}
		close(powCh)
	}()

	for v := range powCh {
		fmt.Println(v)
	}
}
