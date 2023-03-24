package main

import "fmt"

func main() {
	v1 := 0 //int var
	v2 := 0 //int var

	/* Pass by value is when a value of some variable, say X is copied into another such variable, say X1. This copied variable X1 when called into a method is said to be passed by value. Nothing happens
	   to the original variable X. However, in pass by reference, the memory address of this variable will be passed instead of its value. So instead of a duplicate or a copy, the original variable X is only passed.
	   So any changes that happen to the variable in that method are implied on the original variable only. */

	//initializing the values here
	fmt.Println("Memory Address: v1:%p, v2:%p\n", &v1, &v2)
	fmt.Println("Value: v1:%d, v2:%d\n", v1, v2) //were given as 0 each

	//Pass by Value
	Swap(v1, v2) // Golang copies the original value of the two variables v1 and v2
	Add(v1)      // original value of v1
	Add(v2)      // original value of v2

	//Pass by Reference
	Swapref(&v1, &v2) //Golang copies the '&'variable value which is nothing but the memory address initialized above
	Addref(&v1)
	Addref(&v2)

	// printing the final values
	fmt.Println("Demonstration of the Pass by Value VS Pass by Reference procedures:\n")
	fmt.Println("Memory Location: v1:%p, v2:%p\n", &v1, &v2)
	fmt.Println("Value: v1:%d, v2:%d\n", v1, v2)

}

// PASS BY VALUE FUNCTIONS
//Defining the Swap() function
func Swap(x, y int) {
	fmt.Println("Interchange the memory addresses: %p, %p\n", &x, &y)
	x, y = y, x
}

//Defining the Add() function
func Add(x int) {
	fmt.Println("Before adding, memory address was: %p, and Value was: %d\n", &x, x)
	x++
	fmt.Println("After adding, memory address is: %p, and Value is: %d\n", &x, x)
}

//PASS BY REFERENCE FUNCTIONS
/* Since we are working with %p used to denote a pointer,
the datatype being considered will be *int, wherein '*' represents a pointer integer*/

//Defining Swapref() function
func Swapref(x, y *int) {
	fmt.Println("Interchanging memory address through Reference: %p, %p\n", x, y)
	*x, *y = *y, *x
}

//Defining Addref() function
func Addref(x *int) {
	fmt.Println("the Addref Function\n")
}

//path to run in terminal followed with 'go run passby.go'
//C:\Users\hi\OneDrive\Documents\GitHub\Go_lecture\passby.go
