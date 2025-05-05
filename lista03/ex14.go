package main

import "fmt"

func ehPrimo(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n1, n2 int
	fmt.Scanf("%d %d", &n1, &n2)

	for i := n1; i <= n2; i++ {
		if ehPrimo(i) {
			fmt.Println(i)
		}
	}
}
