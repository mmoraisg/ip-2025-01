package main

import (
	"fmt"
	"math"
)

func main() {
	vetorNumeros := make([]float64, 100)
	for i := 0; i < 100; i++ {
		fmt.Scan(&vetorNumeros[i])
	}

	soma := 0.0
	for i := 0; i < 50; i++ {
		diferenca := vetorNumeros[i] - vetorNumeros[99-i]
		soma += math.Pow(diferenca, 3)
	}
	fmt.Printf("O valor da soma e de %.2f.", soma)
}
