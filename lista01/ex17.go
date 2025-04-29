package main
import "fmt"

func main() {
    var x, y int
    fmt.Scanf("%d %d", &x, &y)
    
    if x%2==0 {
        for i:=0;i<y;i++ {
            fmt.Printf("%d ", i*2+x)
        }
    } else {
        fmt.Printf("O NUMERO %d NAO E PAR", x)
    }
}
