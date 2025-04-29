package main
import "fmt"

func main() {
    var a, b, c int
    fmt.Println("DIGITE 3 DIGITOS")
    fmt.Scan(&a)
    fmt.Scan(&b)
    fmt.Scan(&c)
    
    if a>9 || b>9 || c>9 {
        fmt.Println("DIGITO INVALIDO")
    } else {
        k:=a*100+b*10+c
        w:=k*k
        fmt.Printf("%d, %d", k, w)
    }
}
