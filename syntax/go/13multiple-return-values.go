package main 

import "fmt"
/*
	Go has support for multiple return values eg.to return both result and error values.
*/
func vals()(int, int){
	return 3, 7
}

func main(){

	a, b := vals()
	fmt.Println(a,b)

	_, c := vals()
	fmt.Println(c)
}
