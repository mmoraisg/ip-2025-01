package main

import (
	"fmt"
	"math"
)

func fatorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * fatorial(n-1)
}
func main() {
	var x, s float64
	fmt.Scan(&x)

	for i := 0; i < 20; i++ {
		termo := (math.Pow((-1), float64(i))) * x / float64(fatorial(i))
		s += termo
	}
	fmt.Printf("%.2f.", s)
}
