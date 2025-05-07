package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Raio(cm)\t\tVolume da esfera(cm cubicos)\n")
	for i := 0.0; i <= 20; i += 0.5 {
		volume := 4 * math.Pi * math.Pow(i, 3)
		fmt.Printf("%.2f\t\t%.2f\n", i, volume)
	}
}
