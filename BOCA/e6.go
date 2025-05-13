package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	vetor := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&vetor[i])
	}
	sort.Ints(vetor)

	for _, num := range vetor {
		fmt.Print(num)
		fmt.Print(" ")
	}
}
