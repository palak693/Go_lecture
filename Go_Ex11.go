package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ThisorThat   bool = true
	LargestInt uint64 = 1<<32 - 1
	num complex128 = cmplx.Sqrt(25)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ThisorThat, ThisorThat)
  fmt.Printf("Monkeys like hanging on trees? True or False?", ThisorThat)
	fmt.Println("Type: %T Value: %v\n", LargestInt, LargestInt)
	fmt.Printf("Type: %T Value: %v\n", num, num)
}
