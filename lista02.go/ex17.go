package main

import "fmt"

func main() {
	var (
		conta   int
		consumo float64
		tipo    byte
		tipoStr string
	)

	fmt.Print("DIGITE O NUMERO DA CONTA:")
	fmt.Scan(&conta)

	fmt.Print("DIGITE O CONSUMO:")
	fmt.Scan(&consumo)

	fmt.Print("DIGITE O CARACTER DO TIPO DE CONSUMIDOR:")
	fmt.Scan(&tipoStr)
	if len(tipoStr) > 0 {
		tipo = tipoStr[0]
	} else {
		fmt.Println("ERRO: Nenhuma letra foi digitada.")
		return
	}

	if tipo == 'R' {
		gasto := 5 + consumo*0.05
		fmt.Println("CONTA=", conta)
		fmt.Printf("VALOR DA CONTA=%.2f", gasto)
	}
	if tipo == 'C' {
		gasto := 500 + (consumo-80)*0.25
		fmt.Println("CONTA=", conta)
		fmt.Printf("VALOR DA CONTA=%.2f", gasto)
	}
	if tipo == 'I' {
		gasto := 800 + (consumo-100)*0.04
		fmt.Println("CONTA=", conta)
		fmt.Printf("VALOR DA CONTA=%.2f", gasto)
	}
}
