package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	if n < 1 || n > 100 {
		return
	}

	matriz := make([][]int, n)
	for i := 0; i < n; i++ {
		matriz[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&matriz[i][j])
		}
	}

	resultado := make([][]int, n)
	for i := 0; i < n; i++ {
		resultado[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			resultado[j][n-1-i] = matriz[i][j]
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", resultado[i][j])
		}
		fmt.Println()
	}
}
