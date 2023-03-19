// Variables declared without an explicit initial value are given their zero value.
// it is 0 for int, false for bools, "" for strings

package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
  fmt.Println("The only neutral integer on the number line is: ", i)
  fmt.Println("Numbers begin from this integer onwards: ", f)
  fmt.Println("There are 366 days in a non-leap year? True or false?", b)
  fmt.Println("An empty dialogue can look like this-", s)
	fmt.Println("%v %v %v %q\n", i, f, b, s)
}

