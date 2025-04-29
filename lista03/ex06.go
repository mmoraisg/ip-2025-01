package main

import "fmt"

func numTriangular(n int) bool {
	if n <= 0 {
		return false
	}
	for i := 1; i*(i+1)*(i+2) <= n; i++ {
		if i*(i+1)*(i+2) == n {
			return true
		}
	}
	return false
}
func main() {
	var numero int

	for {
		fmt.Print("DIGITE UM NUMERO INTEIRO POSITIVO: ")
		fmt.Scan(&numero)

		if numero > 0 {
			break
		}
		fmt.Println("NUMERO INVALIDO. TENTE NOVAMENTE.")
	}
	if numTriangular(numero) {
		fmt.Printf("%d e um numero triangular.", numero)
	} else {
		fmt.Printf("%d nao e um numero triangular.", numero)
	}
}
