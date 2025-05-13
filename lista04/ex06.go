package main

import "fmt"

func inverter(vetor []int) {
	n := len(vetor)
	for i := 0; i < n/2; i++ {
		vetor[i], vetor[n-1-i] = vetor[n-1-i], vetor[i]
	}
}

func main() {
	vetor := make([]int, 100)
	for i := 0; i < 100; i++ {
		vetor[i] = i + 1
	}
	inverter(vetor)

	for _, invertidos := range vetor {
		fmt.Print(invertidos)
		fmt.Print(" ")
	}

}
