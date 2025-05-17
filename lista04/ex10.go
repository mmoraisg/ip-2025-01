package main

import (
	"fmt"
)

func main() {
	var (
		vetorFibonacci [50]uint64
	)
	vetorFibonacci[0] = 1
	vetorFibonacci[1] = 1

	for i := 2; i < 50; i++ {
		vetorFibonacci[i] = vetorFibonacci[i-1] + vetorFibonacci[i-2]
	}

	for _, num := range vetorFibonacci {
		fmt.Print(num)
		fmt.Print(" ")
	}
}
