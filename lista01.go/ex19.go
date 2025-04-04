package main
import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    
    if n<1 {
        fmt.Println("NUMERO INVALIDO")
    } else {
        var soma float64
        for i:=1;i<=n;i++ {
            soma+=1/float64(i)
        }
        fmt.Printf("%.6f\n", soma)
    } 
}
