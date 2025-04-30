package main

import (
	"fmt"
)

func main() {
	var (
		n, qntdTotal, qntdPar, qntdImpar                                          int
		numero, somaNum, media, maior, menor, somaPar, mediaPar, porcentagemImpar float64
		temValor                                                                  bool = false
	)
	fmt.Println("DIGITE A QUANTIDADE DE NUMERO QUE SERA ENVIADOS.")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Scan(&numero)

		somaNum += numero
		qntdTotal++
		media = somaNum / float64(qntdTotal)

		if !temValor {
			maior = numero
			menor = numero
			temValor = true
		} else {
			if numero > maior {
				maior = numero
			}
			if numero < menor {
				menor = numero
			}
		}

		if int(numero)%2 == 0 {
			qntdPar++
			somaPar += numero
			mediaPar = somaPar / float64(qntdPar)
		} else {
			qntdImpar++
			porcentagemImpar = (float64(qntdImpar) / float64(qntdTotal)) * 100
		}
	}
	fmt.Printf("A soma dos numeros acima e igual a  %.2f.\n", somaNum)
	fmt.Printf("A quantidade de numeros digitados e igual a  %d.\n", qntdTotal)
	fmt.Printf("A media dos numeros digitados e igual a  %.2f.\n", media)
	if temValor {
		fmt.Printf("O maior número digitado foi %.2f.\n", maior)
		fmt.Printf("O menor número digitado foi %.2f.\n", menor)
	}
	fmt.Printf("A media dos numeros pares e de %.2f.\n", mediaPar)
	fmt.Printf("A porcentagem dos numeros impares e de %.2f%%.", porcentagemImpar)
}
