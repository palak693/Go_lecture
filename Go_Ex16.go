package main
import "fmt"
func needInt(x int)int{return x*20 + 2}
func needFloat(x float64)float64{return x*0.2}

func main() {
  fmt.Println(needInt(4))
fmt.Println(needInt(3))
}
