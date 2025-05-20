package main

import (
	"fmt"
)

func main() {
	vetorNumeroDado := make([]int, 4)
	repeticoes := make(map[int]int)

	for i := 0; i < 4; i++ {
		fmt.Scan(&vetorNumeroDado[i])

		for vetorNumeroDado[i] < 1 || vetorNumeroDado[i] > 6 {
			fmt.Println("ERRO: digite numeros entre 1 e 6.")
			fmt.Scan(&vetorNumeroDado[i])
		}
	}

	for _, num := range vetorNumeroDado {
		repeticoes[num]++
	}

	for num, qntd := range repeticoes {
		fmt.Printf("O numero %d aparece %d vezes.\n", num, qntd)
	}
}
