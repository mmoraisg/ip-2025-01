package main

import "fmt"

func main() {
	var a, b, n int

	fmt.Print("DIGITE OS DOIS PRIMEIROS TERMOS DA SERIE DE FETTUCINE (A0 e A1):")
	fmt.Scanf("%d %d", &a, &b)
	fmt.Println("DIGITE O NUMERO DE TERMOS A SEREM IMPRESSOS:")
	fmt.Scan(&n)

	if n < 3 {
		fmt.Println("Numero invalido, n deve ser maior ou igual a 3.")
		return
	}

	fettucine := make([]int, n)
	fettucine[0] = a
	fettucine[1] = b

	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			fettucine[i] = fettucine[i-1] - fettucine[i-2]
		} else {
			fettucine[i] = fettucine[i-1] + fettucine[i-2]
		}
	}
	fmt.Printf("A serie fettucine para n igual a %d:\n", n)

	for i := 0; i <= n; i++ {
		fmt.Printf("%d ", fettucine[i], "\n")
	}
}
