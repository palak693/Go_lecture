package main
import "fmt"
import "strings"

/* the split() method in Go is defined as a "String Library" that breaks the
string into a list ogf substrings by using a specified separator. The split() 
method returns the output of the substring in the form of the slice. */
func main(){

// initializing strings 
  str1 := "Betty, was a little girl, living in a village"
  str2:= "Betty loves to baking cakes"
  str3:= "She likes to eat them too"

//display strings
fmt.Println("String 1: ", str1)
fmt.Println("String 2: ", str2)
fmt.Println("String 3: ", str3)

//splitting the typed out strings using the split() method
r1 := strings.Split(str1, ",")
r2 := strings.Split(str2, "to")
r3 := strings.Split(str3, "")
  
//printing results
  fmt.Println("Res1: ", r1)
  fmt.Println("Res2: ", r2)
  fmt.Println("Res3: ", r3)
}
