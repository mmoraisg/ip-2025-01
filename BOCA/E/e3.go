package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	vetor1 := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&vetor1[i])
	}

	if n == 1 {
		fmt.Print(0)
		return
	}

	for i := 0; i < n; i++ {
		if i == 0 {
			fmt.Printf("%d", vetor1[i+1])
		} else if i == n-1 {
			fmt.Printf("%d", vetor1[i-1])
		} else {
			fmt.Printf("%d", vetor1[i-1]+vetor1[i+1])
		}
		if i < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
