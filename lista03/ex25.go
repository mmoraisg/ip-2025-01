package main

import (
	"fmt"
)

func main() {
	numerador := 16384
	denominador := 0
	soma := 0.0

	for i := 1; i <= 15; i++ {
		denominador = i * i
		if denominador%2 == 0 {
			denominador = denominador * (-1)
		}
		soma += float64(numerador) / float64(denominador)
		numerador /= 2
	}
	fmt.Printf("O valor da soma e de %.2f.\n", soma)
}
