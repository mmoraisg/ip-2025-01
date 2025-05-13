package main

import "fmt"

func main() {
	var (
		n        int
		min, max float64
	)
	fmt.Scan(&n)
	reais := make([]float64, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&reais[i])
		if i == 0 || reais[i] < min {
			min = reais[i]
		}
		if i == 0 || reais[i] > max {
			max = reais[i]
		}
	}

	for i := 0; i < n; i++ {
		if max == min {
			fmt.Printf("0.00")
		} else {
			xnormalizado := (reais[i] - min) / (max - min)
			fmt.Printf("%.2f", xnormalizado)
		}
		if i < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
