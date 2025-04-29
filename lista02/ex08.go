package main
import "fmt"
    
func main() {
    var a float64
    fmt.Print("DIGITE UM NUMERO:")
    fmt.Scan(&a)
    
    if a>20 && a<90 {
        fmt.Printf("%.2f ESTA ENTRE 20 E 90", a)
    } else {
        fmt.Printf("%.2f NAO ESTA ENTRE 20 E 90", a)
    }
}
