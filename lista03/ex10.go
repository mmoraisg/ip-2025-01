package main

import "fmt"

func fibonacci(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		fmt.Print(a, "")
		a, b = b, a+b
	}
}
func main() {
	var n int
	fmt.Print("Digite um numero maior ou igual a 3:")
	fmt.Scan(&n)
	fmt.Print(fibonacci(n))
}
