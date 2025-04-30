package main

import "fmt"

func fatorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else if n < 0 {
		return -1
	}
	return n * fatorial(n-1)
}
func main() {
	var n int
	fmt.Scan(&n)

	resultado := fatorial(n)
	if resultado == -1 {
		fmt.Print("Digite um numero valido.")
	} else {
		fmt.Printf("O fatorial de %d e igual a %d.", n, resultado)
	}
}
