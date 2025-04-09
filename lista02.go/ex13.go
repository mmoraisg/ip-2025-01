package main

import "fmt"

func main() {
    var numero int
    
    for {
        fmt.Scan(&numero)
        dezenas:= (numero/10)%10
        
        if 99>numero || numero>999 {
            break
        }
        fmt.Print(dezenas)
    }
}
