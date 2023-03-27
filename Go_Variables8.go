package main

import "fmt"

// var at package level
var name string
var major string
var grade int 

func main() {
  // var at function level
	var i int
	fmt.Println(i, name, major, grade)
}
