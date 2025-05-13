package main

import (
	"fmt"
	"math"
)

func main() {
	var num int
	vetorRaiz := make([]float64, 15)

	for i := 0; i < 15; i++ {
		fmt.Scan(&num)
		if num == 0 {
			vetorRaiz[i] = -1
		} else {
			vetorRaiz[i] = math.Sqrt(float64(num))
		}
	}
	fmt.Print(vetorRaiz)
}
