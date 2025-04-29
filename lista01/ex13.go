package main
import "fmt"

func main() {
    var nota float64
    fmt.Scan(&nota)
    
    if nota<6 && nota>=0 {
        fmt.Printf("NOTA = %.2f CONCEITO = D", nota)
    }
    if nota>=6 && nota<7.5 {
        fmt.Printf("NOTA = %.2f CONCEITO = C", nota)
    }
    if nota>=7.5 && nota<9 {
        fmt.Printf("NOTA = %.2f CONCEITO = B", nota)
    }
    if nota>=9 && nota<=10 {
        fmt.Printf("NOTA = %.2f CONCEITO = A", nota)
    }
}
