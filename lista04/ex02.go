package main

import "fmt"

func main() {
	var (
		vetor1     [10]int
		vetor2     [5]int
		vetorPar   []int
		vetorImpar []int
	)

	for i := 0; i < 10; i++ {
		fmt.Printf("DIGITE O %dº NUMERO INTEIRO DO VETOR 1:", i+1)
		fmt.Scan(&vetor1[i])
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("DIGITE O %dº NUMERO INTEIRO DO VETOR 2:", i+1)
		fmt.Scan(&vetor2[i])
	}

	soma2 := 0
	for i := 0; i < len(vetor2); i++ {
		soma2 += vetor2[i]
	}

	for _, num := range vetor1 {
		if num%2 == 0 {
			vetorPar = append(vetorPar, num+soma2)
		} else {
			vetorImpar = append(vetorImpar, num+soma2)
		}
	}
	fmt.Println("Primeiro vetor resultante: ", vetorPar)
	fmt.Println("Primeiro vetor resultante: ", vetorImpar)
}
