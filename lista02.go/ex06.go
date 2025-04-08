package main
import "fmt"

func main() {
  var a, b int
  fmt.Scanf("%d %d", &a, &b)
  
  if a%b==0 {
      fmt.Printf("%d e divisivel por %d", a, b)
  } else {
      fmt.Printf("%d nao e divisivel por %d", a, b)
  }
}
