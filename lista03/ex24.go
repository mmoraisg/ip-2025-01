package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("A(rad)\t\tSeno(aproximado)")
	for i := 0.0; i <= 6.3; i += 0.1 {
		fmt.Printf("%.2f\t\t%.2f\n", i, macLaurin(i))
	}
}

func macLaurin(a float64) float64 {
	return a - (math.Pow(a, 3) / 6) + (math.Pow(a, 5) / 120) - (math.Pow(a, 7) / 5040)
}
