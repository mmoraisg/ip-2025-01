package main

import (
	"fmt"
)

func main() {
	vetorIdades := make([]int, 50)
	repeticoes := make(map[int]int)

	for i := 0; i < 50; i++ {
		fmt.Scan(&vetorIdades[i])

		for vetorIdades[i] < 0 {
			fmt.Println("Idade invalida! Digite novamente.")
			fmt.Scan(&vetorIdades[i])
		}
	}

	for _, idade := range vetorIdades {
		repeticoes[idade]++
	}

	var moda int
	maiorFreq := 0

	for idade, qntd := range repeticoes {
		if qntd > maiorFreq {
			maiorFreq = qntd
			moda = idade
		}
	}

	fmt.Printf("A moda das idades e %d (aparece %d vezes).", moda, maiorFreq)
}
