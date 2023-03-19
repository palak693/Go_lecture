// A function can return any number of results, like swap function returning two strings

package main
import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
  a, b := swap("Bhatia", "Reet")
  p, q := swap("First is 2", "Second is 1")
	fmt.Println(a, b)
  fmt.Print("In the number 21,")
  fmt.Println(p, q)
}
