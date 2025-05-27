package main

import (
	"fmt"
)

func main() {
	var n, m, k int
	fmt.Scan(&n, &k, &m)

	if n < 1 || n > 100 || k < 1 || k > 100 || m < 1 || m > 100 {
		return
	}

	matrizA := make([][]int, n)
	for i := 0; i < n; i++ {
		matrizA[i] = make([]int, k)
		for j := 0; j < k; j++ {
			fmt.Scan(&matrizA[i][j])
		}
	}

	matrizB := make([][]int, k)
	for i := 0; i < k; i++ {
		matrizB[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&matrizB[i][j])
		}
	}

	matrizC := make([][]int, n)
	for i := 0; i < n; i++ {
		matrizC[i] = make([]int, m)
		for j := 0; j < m; j++ {
			for t := 0; t < k; t++ {
				matrizC[i][j] += matrizA[i][t] * matrizB[t][j]
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%d ", matrizC[i][j])
		}
		fmt.Println()
	}
}
