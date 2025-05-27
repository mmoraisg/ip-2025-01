package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	vetor := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&vetor[i])
	}
	diferentes := make(map[int]bool)

	for _, num := range vetor {
		diferentes[num] = true
	}
	fmt.Print(len(diferentes))
}
