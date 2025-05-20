package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	vetor := make([]int, 10)

	for i := 0; i < 10; i++ {
		fmt.Scan(&vetor[i])
	}

	if n == 1 {
		fmt.Print(vetor)
	} else if n == 2 {
		sort.Sort(sort.Reverse(sort.IntSlice(vetor)))
		fmt.Print(vetor)
	}
}
