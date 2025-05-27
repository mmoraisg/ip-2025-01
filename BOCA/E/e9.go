package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	vetor1 := make([]int, n)
	vetor2 := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&vetor1[i])
	}
	vetor2[0] = vetor1[0]
	for i := 1; i < n; i++ {
		vetor2[i] = vetor2[i-1] + vetor1[i]
	}
	for _, num := range vetor2 {
		fmt.Print(num)
		fmt.Print(" ")
	}
}
