package main

import (
	"fmt"
)

func main() {
	vetorNotas := make([]float64, 15)
	repeticoes := make(map[float64]int)

	for i := 0; i < 15; i++ {
		fmt.Printf("Digite a nota do aluno %d: ", i+1)
		fmt.Scan(&vetorNotas[i])

		for vetorNotas[i] < 0 || vetorNotas[i] > 10 {
			fmt.Println("Notas invalida. Digite novamente.")
			fmt.Printf("Digite a nota do aluno %d: ", i+1)
			fmt.Scan(&vetorNotas[i])
		}
	}

	fmt.Printf("NOTAS\tFREQUENCIA ABSOLUTA\tFREQUENCIA RELATIVA\n")

	for _, nota := range vetorNotas {
		repeticoes[nota]++
	}

	for nota, freqabsoluta := range repeticoes {
		freqrelativa := float64(freqabsoluta) / 15
		fmt.Printf("%.2f\t\t%d\t\t%.2f\n", nota, freqabsoluta, freqrelativa)
	}
}
