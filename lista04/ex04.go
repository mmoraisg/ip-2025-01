package main

import "fmt"

func main() {
	var vetor [10]int
	repeticoes := make(map[int]int)

	for i := 0; i < 10; i++ {
		fmt.Printf("Digite o elemento %d: ", i+1)
		fmt.Scan(&vetor[i])
		repeticoes[vetor[i]]++
	}
	repetidos := false
	for valor, contagem := range repeticoes {
		if contagem > 1 {
			fmt.Printf("%d se repete %d vezes.\n", valor, contagem)
			repetidos = true
		}
	}

	if !repetidos {
		fmt.Print("Nao ha elementos repetidos no vetor A.")
	}
}
