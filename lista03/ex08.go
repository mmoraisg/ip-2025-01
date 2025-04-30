package main

import "fmt"

func main() {
	var soma int
	for i := 1; i <= 20; i++ {
		fmt.Printf("%d ", i)
		soma += i
	}
	fmt.Print("\n")
	fmt.Printf("A soma dos numeros de 1 a 20 e de %d.\n", soma)
}
