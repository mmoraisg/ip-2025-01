package main

import "fmt"

func main() {
	var b, n int
	fmt.Print("Digite, respectivamente, a base e o expoente:")
	fmt.Scanf("%d %d", &b, &n)

	if b < 2 || n <= 1 {
		fmt.Println("Digite numeros validos (n>1 e b>=2).")
		return
	}

	potencia := 1
	for i := 1; i <= n; i++ {
		potencia *= b
	}
	fmt.Print(potencia)
}
