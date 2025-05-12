package main

import "fmt"

func main() {
	var vetor [10]int

	for i := 0; i < 10; i++ {
		fmt.Printf("DIGITE O %dº NUMERO INTEIRO:", i+1)
		fmt.Scan(&vetor[i])
	}
	encontrado := false

	for i, num := range vetor {
		if num > 50 {
			fmt.Printf("O NUMERO %d DA POSICAO %d É MAIOR QUE 50.\n", num, i+1)
			encontrado = true
		}
	}
	if !encontrado {
		fmt.Println("NAO HA NUMEROS MAIORES QUE 50 NO VETOR.")
	}
}
