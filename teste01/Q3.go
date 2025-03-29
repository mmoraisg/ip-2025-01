package main
import "fmt"

func main() {
    var soma int
   for i:=1;i<101;i++ {
       fmt.Println(i)
       soma +=i
   }
   fmt.Println("A SOMA É = ", soma)
}
