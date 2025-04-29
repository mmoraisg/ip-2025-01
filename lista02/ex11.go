package main
import "fmt"

func main() {
    var x float64
    
    for {
        fmt.Print("DIGITE O VALOR DE X:")
        fmt.Scan(&x)
        
        f:=8/(2-x)
        
        fmt.Println("TENTE OUTRO VALOR PARA A VARIAVEL X.")
        
        if x!=2 {
            fmt.Printf("Para x=%.2f f(%.2f)=%.2f.", x, x, f)
            break
        }
    } 
}
