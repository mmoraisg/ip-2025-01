package main

import "fmt"

func main() {

	vetorAltura := make([]float64, 10)
	for i := 0; i < 10; i++ {
		fmt.Scan(&vetorAltura[i])
	}
	alturas := 0.0
	for _, num := range vetorAltura {
		alturas += num
	}
	media := alturas / (float64(len(vetorAltura)))

	fmt.Print("Alturas acima da media:")

	for _, num := range vetorAltura {
		if num > media {
			fmt.Print(num)
			fmt.Print(" ")
		}
	}
}
