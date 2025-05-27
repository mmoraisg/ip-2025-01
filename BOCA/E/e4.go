package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		n int
		s float64
	)

	fmt.Scan(&n)

	if n%2 != 0 || n < 2 {
		return
	}

	vetor := make([]float64, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&vetor[i])
	}

	for i := 0; i < n/2; i++ {
		dif := vetor[i] - vetor[n-1-i]
		s += math.Pow(dif, 3)
	}
	fmt.Printf("%.2f", s)
}
