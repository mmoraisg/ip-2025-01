package main

import "fmt"

func main() {
	var (
		base     float64
		expoente int
	)
	fmt.Print("DIGITE A BASE E O EXPOENTE DA POTENCIA:")
	fmt.Scan(&base, &expoente)

	potencia := 1.0
	for i := 0; i < expoente; i++ {
		potencia *= base
	}
	fmt.Print(potencia)
}
