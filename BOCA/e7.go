package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	vetorA := make([]int, n)
	vetorB := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&vetorA[i])
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&vetorB[i])
	}

	AB := 0
	for i := 0; i < n; i++ {
		AB += vetorA[i] * vetorB[i]
	}
	fmt.Print(AB)
}
