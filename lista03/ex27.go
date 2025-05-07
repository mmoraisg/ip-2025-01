package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Scan(&x)
	cos := 0.0

	for i := 0; i < 20; i++ {
		exp := 2 * i
		sinal := math.Pow((-1), float64(i))
		numerador := math.Pow(x, float64(i))
		denominador := float64(fatorial(exp))
		termo := sinal * numerador / denominador
		cos += termo
	}
	fmt.Printf("O valor do cosseno pela formula e de %.2f.\n", cos)
	COS := cos - math.Cos(x)
	fmt.Printf("A diferenca entre o valor da formula apresentada em relacao ao cos(%.2f) e de %.2f.\n", x, COS)
}

func fatorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else if n < 0 {
		return -1
	}
	return n * fatorial(n-1)
}
