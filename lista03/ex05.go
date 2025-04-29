<<<<<<< HEAD
package main

import "fmt"

func main() {
	var (
		idade, n     int
		peso, altura float64
	)
	for {
		fmt.Print("IDADE: ")
		fmt.Scan(&idade)

		fmt.Print("PESO: ")
		fmt.Scan(&peso)

		fmt.Print("ALTURA: ")
		fmt.Scan(&altura)

		fmt.Print("DESEJA CONTINUAR DIGITANDOP DADOS (1-sim/ 2- nao)? ")
		fmt.Scan(&n)

		if n == 2 {
			break
		}
	}
}
=======
package main

import "fmt"

func main() {
	var (
		idade, mais50, idade1020, pessoas, menos40 int
		peso, altura, somaAltura                   float64
		resposta                                   string
	)

	for {
		fmt.Println("DIGITE A IDADE:")
		fmt.Scan(&idade)
		fmt.Println("DIGITE A ALTURA:")
		fmt.Scan(&altura)
		fmt.Println("DIGITE O PESO:")
		fmt.Scan(&peso)

		fmt.Println("DESEJA CONTINUAR COLHENDO DADOS DE MAIS PESSOAS? DIGITE sim OU nao. ")
		fmt.Scan(&resposta)

		if idade > 50 {
			mais50++
		}

		if idade >= 10 && idade <= 20 {
			somaAltura += altura
			idade1020++
		}

		pessoas++
		if peso < 40 {
			menos40++
		}

		if resposta == "nao" || resposta == "NAO" {
			break
		}
	}
	var mediaAltura float64
	if idade1020 > 0 {
		mediaAltura = somaAltura / float64(idade1020)
	}
	var porcentagem float64
	porcentagem = (float64(menos40) / float64(pessoas)) * 100

	fmt.Printf("Ha %d pessoas com idade superior a 50 anos.\n", mais50)
	fmt.Printf("A media das alturas das pessoas com idade entre 10 e 20 anos e de %.2f.\n", mediaAltura)
	fmt.Printf("A porcentagem de pessoas com peso inferior a 40kg e de %.2f%%.", porcentagem)
}
>>>>>>> e0fddb2 (Mensagem do commit)
