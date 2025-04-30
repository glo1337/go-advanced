package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	sliceCh := make(chan int)
	powCh := make(chan float64)

	go func() {
		slice := createSlice(10)
		for _, v := range slice {
			sliceCh <- v
		}
		close(sliceCh)
	}()

	go func() {
		for v := range sliceCh {
			powCh <- math.Pow(float64(v), 2)
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
