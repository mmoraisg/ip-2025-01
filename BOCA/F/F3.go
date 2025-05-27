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

	resultado := make([][]int, n)
	for i := 0; i < n; i++ {
		resultado[i] = make([]int, m)
	}

	vizinhos := []struct{ x, y int }{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			soma := 0
			for _, v := range vizinhos {
				ni, nj := i+v.x, j+v.y
				if ni >= 0 && ni < n && nj >= 0 && nj < m {
					soma += matriz[ni][nj]
				}
			}
			resultado[i][j] = soma
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%d ", resultado[i][j])
		}
		fmt.Println()
	}
}
