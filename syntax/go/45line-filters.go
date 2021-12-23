/*
Line filter is common type of prg that reads input
on stdin, processes it and then prints some derived result to stdout.
grep and sed are common lin efilters.
*/

package main 

import(
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){

	scanner := bufio.NewScanner(os.Stdin)
	//wrapping unbuffered os.Stdin with buffered scanner gives us convenient Scan method
	//it advances scanner to next token - which is next line in default scanner.

	for scanner.Scan(){ //Text returns the current token 
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err!=nil{          // check for errors during Scan 
		fmt.Fprintln(os.Stderr, "error:",err)
		os.Exit(1)
	}
}