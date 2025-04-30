package main

import "fmt"

func main() {
	var (
		n, qntdRep, qntdExame, qntdApr int
		n1, n2, media, mediaClasse     float64
	)
	fmt.Print("Informe a quantidade de alunos:")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Printf("Digite as duas notas do aluno %d:", i)
		fmt.Scan(&n1, &n2)

		media = (n1 + n2) / 2.0
		if media <= 3 {
			fmt.Printf("Aluno %d Reprovado.\n", i)
			qntdRep++
		} else if media > 3 && media < 7 {
			fmt.Printf("Aluno %d de Exame.\n", i)
			qntdExame++
		} else {
			fmt.Printf("Aluno %d Aprovado.\n", i)
			qntdApr++
		}
		mediaClasse += media / float64(n)
	}
	fmt.Printf("Total de alunos reprovados:%d.\n", qntdRep)
	fmt.Printf("Total de alunos de exame:%d.\n", qntdExame)
	fmt.Printf("Total de alunos aprovados:%d.\n", qntdApr)
	fmt.Printf("A media da classe e igual a %.2f.\n", mediaClasse)
}
