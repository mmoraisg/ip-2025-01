package main

import (
	"fmt"
)

func primo(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	vetorPrincipal := make([]int, 10)

	for i := 0; i < 10; i++ {
		fmt.Scan(&vetorPrincipal[i])
	}

	for i, num := range vetorPrincipal {
		if primo(num) {
			fmt.Printf("O numero %d, que esta na posicao %d do vetor, e primo.\n", num, i+1)
		}
	}
}
