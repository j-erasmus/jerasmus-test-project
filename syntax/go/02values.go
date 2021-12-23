package main

import "fmt"
import "strconv" // for string conversion

func test() string { 
	//var a string 
	//a =  
	//return "a"+ string(3)  -- cant do this
	return "a"+strconv.Itoa(123) 
}

func main() { 
	fmt.Println(test())
	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
}