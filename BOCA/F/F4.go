package main

import (
	"fmt"
)

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

	linhas := true
	colunas := true

	for i := 0; i < n && linhas; i++ {
		for j := 1; j < m; j++ {
			if matriz[i][j] < matriz[i][j-1] {
				linhas = false
				break
			}
		}
	}

	for j := 0; j < m && colunas; j++ {
		for i := 1; i < n; i++ {
			if matriz[i][j] < matriz[i-1][j] {
				colunas = false
				break
			}
		}
	}

	if linhas && colunas {
		fmt.Println("SIM")
	} else {
		fmt.Println("NAO")
	}
}
