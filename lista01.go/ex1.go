package main
import "fmt"

func main() {
    var a, b, c float64
  fmt.Println("DIGITE AS NOTAS DO ALUNO.")
  fmt.Scanf("%f %f %f", &a, &b, &c)
  
  soma:=a+b+c
  media:=soma/3
  fmt.Printf("MEDIA=%.2f\n", media)
  if media>=6.0 {
      fmt.Println("APROVADO")
  } else {
      fmt.Println("REPROVADO")
  }
}
