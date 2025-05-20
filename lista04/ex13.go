package main

import (
	"fmt"
	"sort"
)

type Empregado struct {
	Numero int
	Meses  int
}

func main() {
	const max = 100
	var empregados []Empregado
	resgistrosUsados := make(map[int]bool)

	for len(empregados) < max {
		var numero, meses int
		fmt.Print("Numero e meses(0 0 para encerrar): ")
		fmt.Scan(&numero, &meses)

		if numero == 0 && meses == 0 {
			break
		}

		if numero <= 0 || meses <= 0 {
			fmt.Println("ERRO: numero e meses devem ser positivos.")
			continue
		}

		if resgistrosUsados[numero] {
			fmt.Println("ERRO: numero de funcionario ja registrado. Digite um numero unico.")
			continue
		}

		resgistrosUsados[numero] = true
		empregados = append(empregados, Empregado{Numero: numero, Meses: meses})
	}

	if len(empregados) < 3 {
		fmt.Println("Numero insuficente de empregados para analise.")
		return
	}

	sort.Slice(empregados, func(i, j int) bool {
		return empregados[i].Meses < empregados[j].Meses
	})

	fmt.Println("\nTres empregados mais recentes:")
	for i := 0; i < 3; i++ {
		fmt.Printf("Empregado %d - %d meses\n", empregados[i].Numero, empregados[i].Meses)
	}
}
