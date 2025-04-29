package main

import "fmt"

func main() {
	var (
		idade, n     int
		peso, altura float64
	)
	for {
		fmt.Print("IDADE: ")
		fmt.Scan(&idade)

		fmt.Print("PESO: ")
		fmt.Scan(&peso)

		fmt.Print("ALTURA: ")
		fmt.Scan(&altura)

		fmt.Print("DESEJA CONTINUAR DIGITANDOP DADOS (1-sim/ 2- nao)? ")
		fmt.Scan(&n)

		if n == 2 {
			break
		}
	}
}
