package main

import "fmt"

func main() {
	var (
		precoI, despesas, deltaP, precoMin, maiorLucro, melhorPreco float64
		qI, deltaQ, melhoringressos                                 int
	)
	fmt.Scanf("%f %d %f %f %d %f", &precoI, &qI, &despesas, &deltaP, &deltaQ, &precoMin)

	fmt.Printf("Preco\tIngressos Vendidos\tLucro\n")

	for i := precoI; i >= precoMin; i -= deltaP {
		lucro := (float64(qI) * precoI) - despesas

		fmt.Printf("%.2f\t\t%d\t\t%.2f\n", precoI, qI, lucro)
		if lucro > maiorLucro {
			maiorLucro = lucro
			melhorPreco = precoI
			melhoringressos = qI
		}
		precoI -= deltaP
		qI += deltaQ
	}
	fmt.Printf("Lucro maximo: %.2f\n", maiorLucro)
	fmt.Printf("Na faixa de preco: %.2f com %d ingressos.", melhorPreco, melhoringressos)
}
