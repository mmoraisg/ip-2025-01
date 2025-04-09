package main
import "fmt"

func main() {
    var compra, venda float64
    
    fmt.Scan(&compra)
    
    if compra<10 {
        venda=1.7*compra
        fmt.Printf("O VALOR DE VENDA DEVE SER: %.2f", venda)
    } else if 10<=compra && compra<30 {
        venda=1.5*compra
        fmt.Printf("O VALOR DE VENDA DEVE SER: %.2f", venda)
    } else if 30<=compra && compra<50 {
        venda=1.4*compra
        fmt.Printf("O VALOR DE VENDA DEVE SER: %.2f", venda)
    } else if compra>=50 {
        venda=1.3*compra
        fmt.Printf("O VALOR DE VENDA DEVE SER: %.2f", venda)
    }
}
