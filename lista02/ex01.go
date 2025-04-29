package main
import "fmt"

func main() {
    var (
        a int
        s string
        )
    fmt.Scan(&a)
    
    switch {
        case a%2==0:
        s="Par"
        case a%2!=0: 
        s="Impar"
    }
    fmt.Print(s)
}
