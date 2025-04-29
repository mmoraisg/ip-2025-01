package main

import (
	"fmt"
)

func main() {
	var matricula, horas int

	fmt.Print("DIGITE A MATRICULA:")
	fmt.Scan(&matricula)
	fmt.Print("DIGITE O NUMERO DE HORA-EXTRA:")
	fmt.Scan(&horas)

	salMin := 788
	valorHE := 10

	salBr := 3*salMin + horas*valorHE
	var salLiq float64

	if salBr > 2000 {
		salLiq = float64(salBr) * 0.8 * 0.88
		fmt.Printf("SALARIO LIQUIDO: R$%.2f", salLiq)
	} else if salBr <= 2000 && salBr > 1500 {
		salLiq = float64(salBr) * 0.88
		fmt.Printf("SALARIO LIQUIDO: R$%.2f", salLiq)
	} else {
		fmt.Printf("SALARIO LIQUIDO: R$%.2f", salBr)
	}
}
