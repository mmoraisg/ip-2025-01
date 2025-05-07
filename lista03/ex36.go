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
	base16 := strconv.FormatInt(int64(n), 16)
	fmt.Printf("O numero %d em binario e %s.\n", n, base16)
}
