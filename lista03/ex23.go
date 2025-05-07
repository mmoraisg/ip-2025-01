package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite o termo ate onde se procedera a soma:")
	fmt.Scan(&n)

	numerador := 1000
	denominador := 1
	s := float64(numerador) / float64(denominador)

	for i := 1; i < n; i++ {
		numerador -= 3
		denominador += 1
		if numerador%2 != 0 {
			numerador = numerador * (-1)
		}
		s += float64(numerador) / float64(denominador)
	}
	fmt.Printf("A soma dos %d termos e igual a %.2f.\n", n, s)
}
