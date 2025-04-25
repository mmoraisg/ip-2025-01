package main

import (
	"fmt"
)

func main() {
	var salCarlos float64
	fmt.Print("DIGITE O SALARIO DE CARLOS:")
	fmt.Scan(&salCarlos)

	valorA := salCarlos
	valorB := salCarlos / 3.0

	meses := 0

	for valorB <= valorA {
		valorA *= 1.02
		valorB *= 1.05
		meses++
	}
	fmt.Printf("A aplicacao de Joao ultrapassara a de Carlos em %d.\n", meses)
	fmt.Printf("Nesta situacao a aplicacao de Carlos sera de R$%.2f, e a de Joao de R$%.2f.", valorA, valorB)
}
