package main
import "fmt"

func main() {
  var f, chuvapol float64
  fmt.Scan(&f)
  fmt.Scan(&chuvapol)
  
  c:=(f-32)*5/9
  chuvamm:=chuvapol*25.4
  
  fmt.Printf("O VALOR EM CELSIUS = %.2f\n", c)
  fmt.Printf("A QUANTIDADE DE CHUVA E = %.2f", chuvamm)
}
