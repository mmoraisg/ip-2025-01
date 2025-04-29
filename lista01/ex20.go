package main
import "fmt"

func main() {
   var h, m, s int
   fmt.Scan(&h)
   fmt.Scan(&m)
   fmt.Scan(&s)
   
   tempo:=s+60*(m+60*h)
   
   fmt.Printf("O TEMPO EM SEGUNDOS E = %d", tempo)
}
