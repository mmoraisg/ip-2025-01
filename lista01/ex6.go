package main
import "fmt"

func main() {
  var n int
    fmt.Scan(&n)  
    
    for i:=0;i<n;i++ { 
        var f float64
        fmt.Scan(&f)
        c:=(f-32)*5/9
        fmt.Printf("%.2f FAHRENHEIT EQUIVALE A %.2f CELSIUS\n", f, c)
    } 
}
