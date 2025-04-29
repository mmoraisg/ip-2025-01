package main
import "fmt"

func main() {
    var salariomin, gastokw float64
    fmt.Println("DIGITE O VALOR DO SALARIO MINIMO")
    fmt.Scan(&salariomin)
    fmt.Println("DIGITE O GASTO EM KW")
    fmt.Scan(&gastokw)
    
    custopkw:=salariomin*7/1000
    custocons:=gastokw*custopkw
    custodesc:=custocons*9/10
    
    fmt.Printf("CUSTO POR KW: R$ %.2f\n", custopkw)
    fmt.Printf("CUSTO DO COSNUMO: R$ %.2f\n",  custocons) 
    fmt.Printf("CUSTO COM DESCONTO: R$ %.2f\n", custodesc)
}
