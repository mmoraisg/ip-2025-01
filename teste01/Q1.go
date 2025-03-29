package main
import "fmt"

func main() {
  var n float64
  fmt.Scan(&n)
  if n>0{
      fmt.Printf("O NÚMERO É POSITIVO")
  } else if n<0 {
      fmt.Printf("O NÚMERO É NEGATIVO")
  } else {
      fmt.Printf("O NÚMERO É NULO")
  }
}
