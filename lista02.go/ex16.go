package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scanf("%f %f %f", &a, &b, &c)

	delta := (b * b) - (4 * (a) * (c))

	if delta < 0 {
		fmt.Print("RAIZES IMAGINARIAS")
		return
	} else {
		if delta != 0 {
			fmt.Println("RAIZES DISTINTAS")
			x1 := ((-1)*(b) + math.Sqrt(delta)) / 2 * a
			x2 := ((-1)*(b) - math.Sqrt(delta)) / 2 * a
			fmt.Print(x1, x2)
		} else {
			fmt.Println("RAIZ UNICA")
			x := ((-1) * (b)) / 2 * a
			fmt.Print(x)
		}
	}

}
