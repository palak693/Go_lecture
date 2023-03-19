package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func subt(a int, b int) int{
return a-b}

/* func div(p int, q int){
return p/q} */


func main() {
	fmt.Println("The sum of 42, 13 is given as",add(42, 13))
  fmt.Println("The difference between 25 and 12 is given as", subt(25, 12))
  // fmt.Println("The quotient obtained between 4 and 2 is given as", div(4, 2))
}
