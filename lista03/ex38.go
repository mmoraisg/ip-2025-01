package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		cpf                        string
		verificador1, verificador2 int
	)
	fmt.Scan((&cpf))

	if len(cpf) != 11 {
		fmt.Print("CPF invalido. Deve conter 11 digitos.")
		return
	}

	cpfSemDigitos := cpf[:9]

	pesos1 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2}
	soma1 := 0

	for i, char := range cpfSemDigitos {
		digito, err := strconv.Atoi(string(char))
		if err != nil {
			fmt.Printf("Caractere invalido na posicao %d.\n", i)
			continue
		}
		produto1 := digito * pesos1[i]
		soma1 += produto1
	}
	resto1 := soma1
	for resto1 >= 11 {
		resto1 -= 11
	}
	if resto1 < 2 {
		verificador1 = 0
	} else if resto1 >= 2 {
		verificador1 = 11 - resto1
	}
	cpf1 := cpfSemDigitos + strconv.Itoa(verificador1)

	pesos2 := []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
	soma2 := 0

	for i, char := range cpf1 {
		digito, err := strconv.Atoi(string(char))
		if err != nil {
			fmt.Printf("Caractere invalido na posicao %d.\n", i)
			continue
		}
		produto2 := digito * pesos2[i]
		soma2 += produto2
	}
	resto2 := soma2
	for resto2 >= 11 {
		resto2 -= 11
	}
	if resto1 < 2 {
		verificador2 = 0
	} else if resto2 >= 2 {
		verificador2 = 11 - resto2
	}
	cpf2 := cpf1 + strconv.Itoa(verificador2)

	if cpf == cpf2 {
		fmt.Print("CPF valido!")
	} else {
		fmt.Println("CPF invalido! Digito(s) verificador(es) incorretos.")
		fmt.Printf("CPF digitado: %s; CPF resultante da verificacao: %s.", cpf, cpf2)
	}
}
