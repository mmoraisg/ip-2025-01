package main

import "fmt"

func main() {
	soma := 0
	for i := 50; i <= 70; i++ {
		if i%2 == 0 {
			soma += i
		}
	}
	fmt.Print(soma)
}
