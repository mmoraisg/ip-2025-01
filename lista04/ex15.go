package main

import (
	"fmt"
)

func main() {
	vetor := make([]int, 30)
	vetorResul := make([]int, 0, 30)

	for i := 0; i < 30; i++ {
		fmt.Scan(&vetor[i])

		if i%2 == 0 {
			vetorResul = append(vetorResul, 2*vetor[i])
		} else {
			vetorResul = append(vetorResul, 3*vetor[i])
		}
	}
	fmt.Print(vetorResul)

}
