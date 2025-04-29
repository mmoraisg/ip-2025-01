package main
import "fmt"

func main() {
    var h int 
    fmt.Scan(&h)
    
    if h%3==0 {
        preco:=h/3*10
        fmt.Printf("O VALOR A PAGAR E = R$%d", preco)
    } 
    if h%3==1 {
        preco:=(h-1)/3*10+5
        fmt.Printf("O VALOR A PAGAR E = R$%d", preco)
    }
    if h%3==2 {
        preco:=(h-2)/3*10+10
        fmt.Printf("O VALOR A PAGAR E = R$%d", preco)
    }
}
