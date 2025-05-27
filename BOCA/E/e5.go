package main

import (
	"fmt"
)

func crescente(vetor []int) int {
	max := 1
	atual := 1

	for i := 1; i < len(vetor); i++ {
		if vetor[i] > vetor[i-1] {
			atual++
			if atual > max {
				max = atual
			}
		} else {
			atual = 1
		}
	}
	return max
}

func main() {
	var (
		n int
	)
	fmt.Scan(&n)

	vetor := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&vetor[i])
	}
	resultado := crescente(vetor)
	fmt.Println(resultado)
}
