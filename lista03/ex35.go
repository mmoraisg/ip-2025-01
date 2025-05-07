package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	_, err := fmt.Scan(&n)

	if err != nil || n < 0 {
		fmt.Print("Numero invalido! Insira um numero inteiro positivo.")
		return
	}
	binario := strconv.FormatInt(int64(n), 2)
	fmt.Printf("O numero %d em binario e %s.\n", n, binario)
}
