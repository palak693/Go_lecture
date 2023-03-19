package main

import "math"
import "fmt"

func main() {
  fmt.Println("The value of Pi is given as: ", math.Pi)
  fmt.Println("The value of Phi function is given as: ", math.Phi)
  // if it was fmt.Println(math.phi), then it would give an error because exports happen to modules or alias with capital beginning letters.
}
