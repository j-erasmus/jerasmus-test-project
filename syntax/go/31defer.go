/*
Defer is used to ensure that a function call is performed later in program's execution,
usually for purposes of cleanup.

eg. we wanted to create a file,write to it and then close when we're done.
*/

package main 

import (
	"fmt"
	"os"
)

func main(){

	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File{
	fmt.Println("creating file")
	f,err := os.Create(p)
	if err != nil{
		panic(err)
	}
	return f
}

func writeFile(f *os.File){
	fmt.Println("writing to file")
	fmt.Fprintln(f,"data")
}

func closeFile(f *os.File){
	fmt.Println("closing")
	f.Close()
}
