package main

import (
	"fmt"
	"math"
)

func softmax(linha []int) []float64 {
	n := len(linha)
	resultado := make([]float64, n)

	max := linha[0]
	for _, val := range linha {
		if val > max {
			max = val
		}
	}

	var somaExp float64
	for i, val := range linha {
		exp := math.Exp(float64(val - max))
		resultado[i] = exp
		somaExp += exp
	}

	for i := range resultado {
		resultado[i] /= somaExp
	}

	return resultado
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	if n < 1 || n > 100 || m < 1 || m > 100 {
		return
	}

	matriz := make([][]int, n)
	for i := 0; i < n; i++ {
		matriz[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&matriz[i][j])
		}
	}

	for i := 0; i < n; i++ {
		resultado := softmax(matriz[i])
		for j := 0; j < m; j++ {
			fmt.Printf("%.6f ", resultado[j])
		}
		fmt.Println()
	}
}
