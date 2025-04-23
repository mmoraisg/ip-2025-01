package main

import "fmt"

func main() {
	meses := []string{
		"", "janeiro", "fevereiro", "março", "abril", "maio", "junho", "julho", "agosto", "setembro", "outubro", "novembro", "dezembro",
	}
	fmt.Println("DIGITE UMA DATA (FORMATO XX/XX/XXXX):")

	var d, m, a int
	fmt.Scanf("%d %d %d", &d, &m, &a)

	if m < 0 || m > 12 {
		fmt.Print("MES INVALIDO.")
		return
	} else if m == 1 || m == 3 || m == 5 || m == 7 || m == 8 || m == 10 || m == 12 {
		if d > 31 || d < 1 {
			fmt.Print("DIA INVALIDO.")
			return
		}
	} else if m == 4 || m == 6 || m == 9 || m == 11 {
		if d < 1 || d > 30 {
			fmt.Print("DIA INVALIDO.")
			return
		}
	}
	if a%4 == 0 && a%100 != 0 || a%4 == 0 && a%100 == 0 && a%400 == 0 {
		if m == 2 {
			if d < 1 || d > 29 {
				fmt.Print("DIA INVALIDO.")
				return
			}
		}
	} else {
		if m == 2 {
			if d < 1 || d > 28 {
				fmt.Print("DIA INVALIDO.")
				return
			}
		}
	}
	fmt.Printf("%d de %s de %d", d, meses[m], a)
}
