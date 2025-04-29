package main

import "fmt"

func main() {
    var idade int
    var classificacao string
    
    fmt.Print("DIGITE A IDADE:") 
    fmt.Scan(&idade)
    
    switch {
    case idade>=0 && idade<=2:
        classificacao= "Recém-nascido"
    case idade>=3 && idade<=11:
        classificacao= "Criança"
    case idade>=12 && idade<=19:
        classificacao= "Adolescente"
    case idade>=20 && idade<=55:
        classificacao= "Adulto"
    case idade>55:
        classificacao= "Idoso"
    default:
        classificacao= "Idade invalida!"
    }
    
    fmt.Println("Classificacao:", classificacao)
}
