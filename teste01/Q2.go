package main
import "fmt"

func main() {
    var n float64
    fmt.Scan(&n)
    if n>20 && n<90 {
        fmt.Printf("O NÚMERO ESTÁ ENTRE 20 E 90")
    } else {
        fmt.Printf("O NÚMERO NÃO ESTÁ ENTRE 20 E 90")
    }
}
