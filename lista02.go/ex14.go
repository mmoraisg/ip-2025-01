package main
import "fmt"

func main() {
    var valoricarro float64
    var a, b, c, d int
    
    fmt.Print("DIGITE O VALOR DO VEICULO:")
    fmt.Scan(&valoricarro)
    fmt.Println("O CARRO POSSUI AS SEGUINTES OPCOES (DIGITE 1 PARA SIM, E 2 PARA NAO): ")
    fmt.Print("Ar condicionado (R$ 1750,00):")
    fmt.Scan(&a)
    fmt.Print("Pintura metalica (R$ 800,00):")
    fmt.Scan(&b)
    fmt.Print("Vidro eletrico (R$ 1200,00):")
    fmt.Scan(&c)
    fmt.Print("Direcao hidraulica (R$ 2000,00):")
    fmt.Scan(&d)
    
    if a==1 {
        valoricarro+=1750
    }
    if b==1 {
        valoricarro+=800
    }
    if c==1 {
        valoricarro+=1200
    }
    if d==1 {
        valoricarro+=2000
    }
    fmt.Printf("VALOR FINAL DO VEICULO R$%.2f", valoricarro)
}
