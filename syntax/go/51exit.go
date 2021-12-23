package main 

import "fmt"
import "os"

func main(){
	defer fmt.Println("!")  // defers will not be run when using os.Exit so this will never be called.

	os.Exit(3)
}