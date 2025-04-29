package main
import "fmt"

func main() {
  var (
      r, h float64
      )
  fmt.Scan(&r)
  fmt.Scan(&h)
  
  pi:=3.14159
  a:=2*pi*r*(r+h)
  custo:=a*100
  
  fmt.Printf("O VALOR DE CUSTO E = R$%.2f", custo)
}
