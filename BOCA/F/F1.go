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

	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			fmt.Print(matriz[i][j])
			if i < n-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
