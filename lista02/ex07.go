package main
import (
    "fmt"
    "sort"
    )

func main() {
    var a, b, c int
   
    for {
         fmt.Print("DIGITE TRES NUMEROS INTEIROS DIFERENTES:")
         fmt.Scanf("%d %d %d", &a, &b, &c)
        
         if a!=b && b!=c && c!=a {
             break
         }
         fmt.Println("ERRO! OS NUMEROS DEVEM SER DIFERENTES. TENTE NOVAMENTE.")
    }
    numeros:=[]int{a, b, c}
    sort.Ints(numeros)
    fmt.Println("NUMEROS EM ORDEM CRESCENTE:", numeros[0], numeros[1], numeros[2])
}
