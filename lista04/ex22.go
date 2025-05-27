package main

import (
	"fmt"
)

func buscarIndice(vetorCodigo []int, codigo int) int {
	for i, cod := range vetorCodigo {
		if cod == codigo {
			return i
		}
	}
	return -1
}

func main() {
	vetorCodigo := make([]int, 10)
	vetorSaldo := make([]float64, 10)
	registrosUsados := make(map[int]bool)

	var codigo int
	var valor float64

	for i := 0; i < 10; {
		fmt.Print("Digite o codigo bancario:")
		fmt.Scan(&codigo)

		if registrosUsados[codigo] {
			fmt.Println("ERRO: o codigo bancario deve ser unico.")
			continue
		}

		vetorCodigo[i] = codigo
		fmt.Print("Digite o saldo bancario:")
		fmt.Scan(&vetorSaldo[i])

		registrosUsados[vetorCodigo[i]] = true
		i++
	}

	for {
		fmt.Printf("\n--- MENU ---\n1.Efetuar deposito\n2.Efetuar saque\n3.Consultar o ativo bancario\n4.Finalizar programa\n")
		var opcao int
		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			fmt.Print("Digite o codigo da conta para deposito:")
			fmt.Scan(&codigo)

			indice := buscarIndice(vetorCodigo, codigo)
			if indice == -1 {
				fmt.Print("Conta nao encontrada.")
				continue
			}

			fmt.Print("Digite o valor a ser depositado:")
			fmt.Scan(&valor)
			if valor > 0 {
				vetorSaldo[indice] += valor
				fmt.Printf("Deposito realizado com sucesso.\nSaldo atual: R$%.2f.\n", vetorSaldo[indice])
			} else {
				fmt.Println("Valor invalido.")
			}

		case 2:
			fmt.Print("Digite o codigo da conta para deposito:")
			fmt.Scan(&codigo)

			indice := buscarIndice(vetorCodigo, codigo)
			if indice == -1 {
				fmt.Print("Conta nao encontrada.")
				continue
			}

			fmt.Print("Digite o valor a ser sacado:")
			fmt.Scan(&valor)

			if valor > 0 && vetorSaldo[indice] >= valor {
				vetorSaldo[indice] -= valor
				fmt.Printf("Saque realizado com sucesso.\nSaldo atual: R$%.2f.\n", vetorSaldo[indice])
			} else {
				fmt.Print("Saldo insuficiente ou valor invalido.")
			}

		case 3:
			total := 0.0
			for _, saldo := range vetorSaldo {
				total += saldo
			}
			fmt.Printf("Ativo bancario e de R$%.2f.", total)

		case 4:
			fmt.Println("Programa finalizado.")
			return

		default:
			fmt.Println("Opcao invalida.")
		}
	}
}
