package main

import (
	"fmt"
)

func main() {
	var (
		vetor           [10]int
		vetorPar        []int
		vetorImpar      []int
		quantidadeImpar int
	)

	for i := 0; i < 10; i++ {
		fmt.Printf("Digite o elemento %d: ", i+1)
		fmt.Scan(&vetor[i])
	}

	somaPar := 0
	for _, num := range vetor {
		if num%2 == 0 {
			vetorPar = append(vetorPar, num)
			somaPar += num
		} else {
			vetorImpar = append(vetorImpar, num)
			quantidadeImpar++
		}
	}
	fmt.Println("Numeros pares digitados: ", vetorPar)
	fmt.Println("Soma dos numeros pares digitados: ", somaPar)
	fmt.Println("Numeros impares digitados: ", vetorImpar)
	fmt.Println("Quantidade de numeros impares digitados: ", quantidadeImpar)
}
