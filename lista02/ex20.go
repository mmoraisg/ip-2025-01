package main

import (
	"fmt"
)

func main() {
	var codigo int
	var precoNormal float64

	fmt.Print("DIGITE O CODIGO:")
	fmt.Scan(&codigo)

	fmt.Print("DIGITE O PRECO DA ETIQUETA:")
	fmt.Scan(&precoNormal)

	if codigo == 1 {
		preco := 0.9 * precoNormal
		fmt.Printf("O VALOR A PAGAR E R$%.2f.", preco)
	} else if codigo == 2 {
		preco := 0.95 * precoNormal
		fmt.Printf("O VALOR A PAGAR E R$%.2f.", preco)
	} else if codigo == 3 {
		preco := precoNormal
		fmt.Printf("O VALOR A PAGAR E R$%.2f, EM DUAS VEZES.", preco)
	} else if codigo == 4 {
		preco := 1.1 * precoNormal
		fmt.Printf("O VALOR A PAGAR E R$%.2f, EM 3 VEZES.", preco)
	} else {
		fmt.Print("DIGITE UM CODIGO VALIDO.")
		return
	}
}
