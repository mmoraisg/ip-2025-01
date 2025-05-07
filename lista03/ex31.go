package main

import (
	"fmt"
	"math/big"
)

func main() {
	resultado := big.NewInt(1)
	for i := 1; i <= 64; i++ {
		resultado.Mul(resultado, big.NewInt(2))
	}
	fmt.Print(resultado)
}
