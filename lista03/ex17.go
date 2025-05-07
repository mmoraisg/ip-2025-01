package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		for m := 1; m <= 10; m++ {
			fmt.Printf("Indice [%d][%d]\n", i, m)
		}
	}
}
