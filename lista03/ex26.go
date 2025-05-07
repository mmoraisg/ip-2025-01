package main

import (
	"fmt"
)

func main() {
	soma := 0.0

	for i := 0; i < 20; i++ {
		numerador := 100 - i
		denominador := fatorial(i)
		soma += float64(numerador) / float64(denominador)
	}
	fmt.Print(soma)
}

func fatorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else if n < 0 {
		return -1
	}
	return n * fatorial(n-1)
}
