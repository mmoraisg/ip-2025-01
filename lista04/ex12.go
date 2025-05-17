package main

import (
	"fmt"
)

func main() {
	vetorNotas := make([]float64, 15)
	for i := 0; i < 15; i++ {
		fmt.Printf("Digite a nota do aluno %d: ", i+1)
		fmt.Scan(&vetorNotas[i])

		for vetorNotas[i] < 0 || vetorNotas[i] > 10 {
			fmt.Println("Notas invalida. Digite novamente.")
			fmt.Printf("Digite a nota do aluno %d: ", i+1)
			fmt.Scan(&vetorNotas[i])
		}
	}

	for _, num := range vetorNotas {
		fmt.Print(num)
		fmt.Print(" ")
	}
}
