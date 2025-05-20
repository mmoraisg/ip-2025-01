package main

import (
	"fmt"
	"sort"
)

func main() {
	vetor10 := make([]int, 10)
	vetor5 := make([]int, 5)

	fmt.Println("Digite os 10 elementos inteiros do vetor 1.")
	for i := 0; i < 10; i++ {
		fmt.Scan(&vetor10[i])
	}

	fmt.Println("Digite os 5 elementos inteiros do vetor 2.")
	for i := 0; i < 5; i++ {
		fmt.Scan(&vetor5[i])
	}

	sort.Ints(vetor10)

	for _, num := range vetor10 {
		divisiveis := []string{}

		for i, div := range vetor5 {
			if div != 0 && num%div == 0 {
				divisiveis = append(divisiveis, fmt.Sprintf("\tDivisivel por %d, na posicao %d.", div, i+1))
			}
		}

		if len(divisiveis) > 0 {
			fmt.Printf("Numero %d:\n", num)
			for _, msg := range divisiveis {
				fmt.Println(msg)
			}
		}
	}
}
