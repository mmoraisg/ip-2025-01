package main
import "fmt"

func main() {
   var n int
   fmt.Scan(&n)
   
   if n%5==0 {
       fmt.Print("Divisivel por 5")
   } else {
       fmt.Print("Nao divisivel por 5")
   }
}
