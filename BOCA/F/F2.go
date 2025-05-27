package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	matriz := make([][]int, n)

	if n < 1 || m < 1 || n > 100 || m > 100 {
		return
	}

	for i := 0; i < n; i++ {
		matriz[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&matriz[i][j])
		}
	}

	somaLinhas := make([]int, n)
	for i := 0; i < n; i++ {
		soma := 0
		for j := 0; j < m; j++ {
			soma += matriz[i][j]
		}
		somaLinhas[i] = soma
	}

	somaColunas := make([]int, m)
	for j := 0; j < m; j++ {
		soma := 0
		for i := 0; i < n; i++ {
			soma += matriz[i][j]
		}
		somaColunas[j] = soma
	}

	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(somaLinhas[i])
	}
	fmt.Println()

	for j := 0; j < m; j++ {
		if j > 0 {
			fmt.Print(" ")
		}
		fmt.Print(somaColunas[j])
	}
	fmt.Println()
}
