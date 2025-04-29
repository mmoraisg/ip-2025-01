package main

import (
	"fmt"
)

func main() {
	var (
		precoNormal float64
		tipo        string
		diaSema     int
	)

	fmt.Print("DIGITE O PRECO NORMAL DO DVD: R$")
	fmt.Scan(&precoNormal)

	fmt.Print("DIGITE O DIA DA LOCACAO (1 para domingo ao 7 para segunda):")
	fmt.Scan(&diaSema)

	fmt.Print("DIGITE A CATEGORIA DO DVD (lancamento/ comum):")
	fmt.Scan(&tipo)

	if diaSema >= 1 && diaSema <= 7 {
		if diaSema == 2 || diaSema == 3 || diaSema == 5 {
			custo := precoNormal * 0.6
			if tipo == "lancamento" {
				custo *= 1.15
				fmt.Printf("O CUSTO DA LOCACAO E DE R$%.2f", custo)
			} else if tipo == "comum" {
				fmt.Printf("O CUSTO DA LOCACAO E DE R$%.2f", custo)
			}
		} else {
			custo := precoNormal
			if tipo == "lancamento" {
				custo *= 1.15
				fmt.Printf("O CUSTO DA LOCACAO E DE R$%.2f", custo)
			} else if tipo == "comum" {
				fmt.Printf("O CUSTO DA LOCACAO E DE R$%.2f", custo)
			}
		}
	}
}
