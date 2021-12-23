/*
	Go supports anonumous functions which can form closures.
	Rhese are useful when you want to define a function inline without having to name it.
*/

package main 

import "fmt"

func intSeq() func() int{
	i :=1
	return func() int{
		i+=1
		return i
	}
}

func main(){

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}