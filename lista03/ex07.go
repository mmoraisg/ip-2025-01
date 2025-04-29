package main

import "fmt"

func main() {
	var n, qntdTotal, qntdImpar int
	var numero, soma, media float64
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Scan(&numero)

		soma += numero
		qntdTotal++

		if numero%2 != 0 {
			qntdImpar++
		}

	}
}
