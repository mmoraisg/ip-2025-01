package main
import "fmt"

func main() {
    var a, r, n int
    fmt.Scanf("%d %d %d", &a, &r, &n)
    
    soma:=(2*a+r*(n-1))*n/2
    
    fmt.Println(soma)
}
