package main

import (
	"fmt"
)

func main() {
	var n1, n2 int
	fmt.Scanf("%d %d", &n1, &n2)

	if n1 <= 0 || n2 <= 0 {
		fmt.Print("Por favor, insira apenas números inteiros positivos.")
		return
	}

	resultado := mmc(n1, n2)
	fmt.Printf("O MMC de %d e %d e de %d.\n", n1, n2, resultado)
}
func mdc(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
func mmc(a, b int) int {
	return (a * b) / mdc(a, b)
}
