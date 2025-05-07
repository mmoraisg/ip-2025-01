package main

import (
	"fmt"
)

func main() {
	var (
		id, idMaisGordo, idMaisMagro       int
		peso, pesoMaisGordo, pesoMaisMagro float64
	)

	pesoMaisGordo = -1.0
	pesoMaisMagro = 1e9

	for i := 1; i <= 5; i++ {
		fmt.Printf("Digite o número de identificação do boi %d: ", i)
		fmt.Scan(&id)

		fmt.Printf("Digite o peso do boi %d (em kg): ", i)
		fmt.Scan(&peso)

		if peso > pesoMaisGordo {
			pesoMaisGordo = peso
			idMaisGordo = id
		}
		if peso < pesoMaisMagro {
			pesoMaisMagro = peso
			idMaisMagro = id
		}
	}

	fmt.Printf("Boi mais gordo: ID %d com %.2f kg\n", idMaisGordo, pesoMaisGordo)
	fmt.Printf("Boi mais magro: ID %d com %.2f kg\n", idMaisMagro, pesoMaisMagro)
}
