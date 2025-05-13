package main

import (
	"fmt"
	"reflect"
)

func inverter(vetor []int) {
	n := len(vetor)
	for i := 0; i < n/2; i++ {
		vetor[i], vetor[n-1-i] = vetor[n-1-i], vetor[i]
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	vetor := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&vetor[i])
	}

	original := make([]int, len(vetor))
	copy(original, vetor)

	inverter(vetor)

	if reflect.DeepEqual(vetor, original) {
		fmt.Print("SIM")
	} else {
		fmt.Print("NAO")
	}
}
