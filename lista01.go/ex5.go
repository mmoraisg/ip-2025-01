package main

import "fmt"

func main() {
    var(
        conta int
        consumo float64
        tipo byte)
    
    fmt.Print("DIGITE O NUMERO DA CONTA:")
    fmt.Scan(&conta)
    
    fmt.Print("DIGITE O CONSUMO:")
    fmt.Scan(&consumo)
    
    fmt.Print("DIGITE O CARACTER DO TIPO DE CONSUMIDOR:")
    fmt.Scanf("%c", &tipo)
    
    if tipo=='R' {
        gasto:=5+consumo*0.05
        fmt.Println("CONTA=", conta)
        fmt.Printf("VALOR DA CONTA=%.2f\n", gasto)
    }
    if tipo=='C' {
        gasto:=500+(consumo-80)*0.25
        fmt.Println("CONTA=", conta)
        fmt.Printf("VALOR DA CONTA=%.2f\n", gasto)
    }
    if tipo=='I' {
        gasto:=800+(consumo-100)*0.04
        fmt.Println("CONTA=", conta)
        fmt.Printf("VALOR DA CONTA=%.2f\n", gasto)
    }
}
