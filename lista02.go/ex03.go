package main
import "fmt"

func main() {
    var a, b int
    fmt.Scanf("%d %d", &a, &b)
    
    soma:=a+b
    
    if soma>20 {
        somaa:=soma+8
        fmt.Println(somaa)
    } else {
        somaa:=soma-5
        fmt.Println(somaa)
    }
}
