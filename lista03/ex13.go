package main

import "fmt"

func main() {
	numerador := 1.0
	denominador := 1.0
	h := 0.0

	for denominador <= 50 {
		h += numerador / denominador
		numerador += 2
		denominador += 1
	}
	fmt.Printf("O valor de h para a formula acima e de %.2f.", h)
}
