package main
import "fmt"

func main() {
    var salario float64
    fmt.Scan(&salario)
    
    if salario<=300 {
        salariorea:=1.5*salario
        fmt.Printf("SALARIO COM REAJUSTE = %.2f", salariorea)
    } else {
        salariorea:=1.3*salario
        fmt.Printf("SALARIO COM REAJUSTE = %.2f", salariorea)
    }
}
