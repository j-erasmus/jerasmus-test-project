package main 

import "os"
import "fmt"

func main(){

	argsWithPrg := os.Args  //gives access to raw xommand-line arguments.
	argsWithoutProg := os.Args[1:] //first value in slice is path to prg
									// os.Args[1:]holds arguments to the prg

	args := os.Args[3] // can access individual args

	fmt.Println(argsWithPrg)
	fmt.Println(argsWithoutProg)
	fmt.Println(args)
}