package main

import (
	"fmt"
)

func main() {
	var n int

	for {
		fmt.Scan(&n)
		quadrado := false

		for i := 1; i*i <= n; i++ {
			if n == i*i {
				quadrado = true
			}
		}
		if quadrado {
			fmt.Print("Quadrado perfeito.\n")
		} else {
			fmt.Print("Nao e quadrado perfeito.\n")
		}

		if n <= 0 {
			break
		}
	}
}
