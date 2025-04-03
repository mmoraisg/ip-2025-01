package main
import "fmt"

func main() {
    var(
        n, np int
        pop, ger, arq, cad, renda float64)
    fmt.Scan(&n)
    
    for i:=0;i<n;i++ {
        fmt.Scanf("%d %f %f %f %f", &np, &pop, &ger, &arq, &cad)
    }
    for i:=0;i<n;i++ {
    renda=float64(np)*(pop+ger*5+arq*10+cad*20)/ 100
        
        fmt.Printf("A RENDA NO JOGO. %d  E = %.2f\n", i+1, renda)
      }
  }
