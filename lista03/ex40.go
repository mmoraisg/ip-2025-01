package main

import (
	"fmt"
)

func main() {
	var (
		maiorLucro, melhorPreco float64
		melhoringressos         int
	)
	preco := 6.0
	ingressos := 130

	fmt.Printf("Ingressos vendidos\tPreco do ingresso\tLucro\n")
	for i := 6.0; i > 1.0; i -= 0.6 {
		lucro := (float64(ingressos) * preco) - 300

		fmt.Printf("%d\t\t\t%.2f\t\t\t%.2f\n", ingressos, preco, lucro)

		if lucro > maiorLucro {
			maiorLucro = lucro
			melhorPreco = preco
			melhoringressos = ingressos
		}

		preco -= 0.6
		ingressos += 30
	}
	fmt.Printf("O lucro maximo esperado foi de R$%.2f, quando o preco do ingresso foi R$%.2f,\n", maiorLucro, melhorPreco)
	fmt.Printf("assim, vendendo %d ingressos.", melhoringressos)
}
