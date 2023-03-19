package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 8, 4
	var f float64 = math.Sqrt(float64(x*y + y*x))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
