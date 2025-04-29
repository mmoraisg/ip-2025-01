package main

import (
	"fmt"
)

func main() {
	var (
		matricula           int
		n1, n2, n3, mediaEx float64
	)
	fmt.Print("DIGITE A MATRICULA DO ALUNO:")
	fmt.Scan(&matricula)

	fmt.Print("DIGITE AS NOTAS DAS TRES PROVAS DO ALUNO (n1, n2 e n3):")
	fmt.Scan(&n1, &n2, &n3)

	fmt.Print("DIGITE A MEDIA DOS EXERCICIOS DO ALUNO:")
	fmt.Scan(&mediaEx)

	fmt.Printf("Matricula: %d.\n", matricula)
	fmt.Printf("As notas do aluno sao %.2f, %.2f, %.2f.\n", n1, n2, n3)
	fmt.Printf("Media dos exercicios: %.2f.\n", mediaEx)

	mediaFinal := (n1 + n2*2 + n3*3 + mediaEx) / 7.0
	fmt.Printf("Media final: %.2f.\n", mediaFinal)

	switch {
	case mediaFinal <= 10 && mediaFinal >= 9:
		fmt.Print("A APROVADO.")
	case mediaFinal < 9 && mediaFinal >= 7.5:
		fmt.Print("B APROVADO.")
	case mediaFinal < 7.5 && mediaFinal >= 6:
		fmt.Print("C APROVADO.")
	case mediaFinal < 6 && mediaFinal >= 4:
		fmt.Print("D REPROVADO.")
	case mediaFinal < 4:
		fmt.Print("E REPROVADO.")
	}
}
