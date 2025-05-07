package main

import "fmt"

func main() {
	s := 0.0
	denominador := 1
	a := 37
	b := 38
	for denominador <= 37 {
		numerador := a * b
		s += float64(numerador) / float64(denominador)
		denominador += 1
		a -= 1
		b -= 1
	}
	fmt.Print(s)
}
