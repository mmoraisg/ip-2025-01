package main

import "fmt"

func main() {
	var n, q int
	fmt.Scan(&n)
	vetorN := make([]int, n)

	if n < 1 || n > 1000 {
		return
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&vetorN[i])
	}

	fmt.Scan(&q)
	vetorQ := make([]int, q)

	if q < 1 || q > 1000 {
		return
	}

	for i := 0; i < q; i++ {
		fmt.Scan(&vetorQ[i])
	}

	mapa := make(map[int]bool)
	for _, num := range vetorN {
		mapa[num] = true
	}

	for _, consulta := range vetorQ {
		if mapa[consulta] {
			fmt.Println("SIM")
		} else {
			fmt.Println("NAO")
		}
	}
}
