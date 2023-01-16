package main

import (
	"fmt"
)

// first init function will execute then main() function will execute
func init() {
	fmt.Println("in init function")
}

func main() {
	a := 10
	var b string = "mitesh"
	var c float64 = 22.33
	var d bool = true

	// %d,%v  --> integer
	// %s, %v ---> strings

	fmt.Println(a, b)
	fmt.Printf("value of A is %d  %v and datatype is %T\n", a, a, a)
	fmt.Printf("value of B is %s  %v and datatype is %T\n", b, b, b)
	fmt.Printf("value of c is %f %v and datatype is %T\n", c, c, c)
	fmt.Printf("value of D is %t %v and datatype is %T\n", d, d, d)

}
