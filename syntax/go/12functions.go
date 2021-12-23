package main 

import "fmt"

func plus(a int,b int)int {
	return a+b
}

func plusplus(a,b,c int)int{
	return a+b+c
}
func main(){
	fmt.Println("sum:",plus(2,3))
	fmt.Println("sumPlusPlus:",plusplus(2,3,5))
}