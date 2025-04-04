package main
import (
    "fmt"
    "math"
)
func main() {
    var h, a float64
    fmt.Scanf("%f %f", &h, &a)
    
    raiz:=math.Sqrt(3)
    volume:=(a*a*h*raiz)/2
    
    fmt.Printf("O VOLUME DA PIRAMIDE E = %.2f METROS CUBICOS ", volume)
}
