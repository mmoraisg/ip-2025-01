package main

import "fmt"

func main() {
	var vetor [10]int
	repeticoes := make(map[int]int)
	var menor int
	indiceMenor := 0

	fmt.Printf("Digite 10 elementos diferentes entre si!\n")

	for i := 0; i < 10; i++ {
		fmt.Printf("Digite o elemento %d: ", i+1)
		fmt.Scan(&vetor[i])
		repeticoes[vetor[i]]++

		if i == 0 {
			menor = vetor[i]
			indiceMenor = i + 1
		} else if vetor[i] < menor {
			menor = vetor[i]
			indiceMenor = i + 1
		}
	}
	repetidos := false
	for valor, contagem := range repeticoes {
		if contagem > 1 {
			fmt.Printf("%d se repete %d vezes.\n", valor, contagem)
			repetidos = true
			break
		}
	}
	if !repetidos {
		fmt.Printf("O menor elemento do vetor %d, na posicao %d.", menor, indiceMenor)
	}
}
