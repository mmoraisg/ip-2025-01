package main

import (
	"fmt"
)

func main() {
	var n1, n2 int
	fmt.Scanf("%d %d", &n1, &n2)

	n := n1
	for i := 1; i < n2; i++ {
		n += n1
	}
	fmt.Print(n)
}
