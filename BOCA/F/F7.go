package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	if 1 > n || n > 100 || 1 > m || m > 100 {
		return
	}

	matrizEntrada := make([][]int, n)
	for i := 0; i < n; i++ {
		matrizEntrada[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&matrizEntrada[i][j])
		}
	}

	for i := 0; i < n; i++ {
		maior := matrizEntrada[i][0]
		for j := 0; j < m; j++ {
			if matrizEntrada[i][j] > maior {
				maior = matrizEntrada[i][j]
			}
		}

		if maior == 0 {
			maior = 1
		}

		for j := 0; j < m; j++ {
			valor := float64(matrizEntrada[i][j]) / float64(maior)
			fmt.Printf("%.6f ", valor)
		}
		fmt.Println()
	}
}
