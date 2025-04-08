package main
import "fmt"

func main() {
    var n int 
    fmt.Scan(&n)
    
    if n==0 {
        fmt.Print("Nulo")
    } else if n>0 {
        fmt.Print("Positivo")
    } else if n<0 {
        fmt.Print("Negativo")
    }
}
