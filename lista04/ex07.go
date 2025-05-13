package main

import "fmt"

func main() {
	vetor := make([]int, 50)
	indice := 0

	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			vetor[indice] = i
			indice++
		}
	}
	fmt.Print(vetor)
}
