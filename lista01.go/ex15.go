package main
import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    
    if 5>=n || n>=2000 {
        fmt.Print("NUMERO INVALIDO. TENTE UM NUMERO ENTRE 5 E 2000")
    } else {
        for i:=0;i<n;i++ {
            if (i+1)%2==0 {
                fmt.Printf("%dˆ2=%d\n", i+1, (i+1)*(i+1))
            }
        }
    }
}
