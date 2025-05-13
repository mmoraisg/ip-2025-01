package main

import "fmt"

func main() {
	var (
		n                   int
		altura, media, soma float64
		alturas             []float64
	)
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Scan(&altura)
		alturas = append(alturas, altura)
		soma += altura
	}
	media = soma / float64(n)

	for _, a := range alturas {
		if a > media {
			fmt.Println(a)
		}
	}
}
