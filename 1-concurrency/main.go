package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numCount := 10
	sliceCh := make(chan int, numCount)
	powCh := make(chan int, numCount)

	go func() {
		slice := createSlice(numCount)
		for _, v := range slice {
			sliceCh <- v
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

func createSlice(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = rand.Intn(101)
	}

	return slice
}
