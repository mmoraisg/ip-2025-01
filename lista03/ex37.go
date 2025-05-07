package main

import (
	"fmt"
	"strconv"
)

func main() {
	var entrada string
	_, err := fmt.Scan(&entrada)

	n, convErr := strconv.ParseInt(entrada, 8, 64)

	if err != nil || convErr != nil || n < 0 {
		fmt.Print("Numero invalido! Insira um numero inteiro positivo.")
		return
	}
	fmt.Printf("O numero %s em decimal e %d.\n", entrada, n)
}
