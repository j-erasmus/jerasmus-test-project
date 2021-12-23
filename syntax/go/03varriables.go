package main

import "fmt"

func main(){ 
	var a string = "tushar"
	fmt.Println(a)

	var b,c int = 1,2 //declare more than one varriable at a time
	fmt.Println(b,c)

	var d = true    // go will infere the type of d from its value
	fmt.Printf("%T %v \n",d,d)  // printing type using %T

	e := 3   //shorthand operator for declaring and initialization
	fmt.Println(e)
}