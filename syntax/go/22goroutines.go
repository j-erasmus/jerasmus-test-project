/*
goroutine is a lightweight thread of execution.

*/

package main 

import "fmt"

func f(from string){
	for i:=0;i<3;i++{
		fmt.Println(from,":",i)
	}
}

func main(){

	f("direct")   // usual way of calling fn , synchronously

	go f("goroutine")  // invoke f() in a goroutine - execute concurrently with calling fn

	go func(msg string){   // goroutine for a anonymous function call.
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}