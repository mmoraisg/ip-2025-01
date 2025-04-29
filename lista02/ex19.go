package main

import (
	"fmt"
	"math"
)

func main() {
	var forma string
	var h, r float64

	fmt.Print("DIGITE A OPCAO DE FIGURA DESEJADA (cone/ cilindro/ esfera):")
	fmt.Scan(&forma)

	if forma == "cone" {
		fmt.Print("DIGITE O RAIO E A ALTURA DO CONE:")
		fmt.Scan(&r, &h)

		volume := (math.Pi * r * r * h) / 3
		area := math.Pi * r * math.Sqrt(r*r+h*h)

		fmt.Printf("O VOLUME DO CONE ESCOLHIDO E DE %.2f u.v. JA A AREA E DE %.2f u.a.\n", volume, area)

	} else if forma == "cilindro" {
		fmt.Print("DIGITE O RAIO E A ALTURA DO CILINDRO:")
		fmt.Scan(&r, &h)

		volume := math.Pi * r * r * h
		area := 2 * math.Pi * r * h

		fmt.Printf("O VOLUME DO CILINDRO ESCOLHIDO E DE %.2f u.v. JA A AREA E DE %.2f u.a.\n", volume, area)

	} else if forma == "esfera" {
		fmt.Print("DIGITE O RAIO DA ESFERA:")
		fmt.Scan(&r)

		volume := (4.0 / 3.0) * math.Pi * r * r * r
		area := 4 * math.Pi * r * r

		fmt.Printf("O VOLUME DA ESFERA ESCOLHIDA E DE %.2f u.v. JA A AREA E DE %.2f u.a.\n", volume, area)

	} else {
		fmt.Print("ERRO: Escreva uma forma geometrica dentre as opcoes acima.")
		return
	}
}
