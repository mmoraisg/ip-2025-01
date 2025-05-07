package main

import (
	"fmt"
	"math"
)

func main() {
	soma := 0.0
	sinal := 1.0

	for i := 0; i < 51; i++ {
		denominador := 2*i + 1
		termo := sinal / math.Pow(float64(denominador), 3)
		soma += termo
		sinal *= -1
	}
	pi := math.Cbrt(32 * soma)
	fmt.Print(pi)
}
