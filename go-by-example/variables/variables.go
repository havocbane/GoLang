package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2 // multiple variables declared and initialized at once
	fmt.Println(b, c)

	var d = true // type inference
	fmt.Println(d)

	var e int // uninitialized variables are zero-valued
	fmt.Println(e)

	f := "apple" // syntactic sugar for declaring and initializing a variable (only available inside functions)
	fmt.Println(f)
}
