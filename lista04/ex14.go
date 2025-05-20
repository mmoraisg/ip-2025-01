package main

import (
	"fmt"
)

func main() {
	vetor1 := make([]int, 10)
	vetor2 := make([]int, 10)

	for i := 0; i < 10; i++ {
		fmt.Printf("Digite um valor inteiro para o termo %d do vetor 1:", i+1)
		fmt.Scan(&vetor1[i])
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("Digite um valor inteiro para o termo %d do vetor 2:", i+1)
		fmt.Scan(&vetor2[i])
	}

	vetorRes := make([]int, 0, 20)

	for i := 0; i < 10; i++ {
		vetorRes = append(vetorRes, vetor1[i], vetor2[i])
	}
	fmt.Println("Vetor resultante intercalado:", vetorRes)
}
