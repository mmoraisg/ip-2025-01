package main

import (
	"fmt"
)

func main() {
	var n1, n2, quociente int
	fmt.Scanf("%d %d", &n1, &n2)
	n := n1
	for n >= n2 {
		n -= n2
		quociente++
	}
	fmt.Printf("O resto da divisao de %d por %d e de %d.\n", n1, n2, n)
	fmt.Printf("Ja o quociente desta divisao e de %d.", quociente)
}
