/*
Go offers support for string formatting in the printf traditon.
*/

package main 

import(
	"fmt"
	"os"
)

type point struct {
	x,y int
}

func main(){
	var p point
	p.x = 1
	p.y =2

	fmt.Printf("%v \n", p)

	fmt.Printf("%+v\n",p)

	fmt.Printf("%#v\n",p)

	fmt.Printf("%T\n",p)

	fmt.Printf("%t\n",true)

	fmt.Printf("%d\n",123)

	fmt.Printf("%c\n",97) //character of the given value

	fmt.Printf("%x\n",456) //hex value

	fmt.Printf("%f\n",78.9) //floats

	fmt.Printf("%e\n",12340000.0) //scientific version

	fmt.Printf("%E\n",12340000.0) // --''--

	fmt.Printf("%q\n","\"string\"") //double quote strings use %q

	fmt.Printf("%x\n","hex this")
	fmt.Printf("%p\n",&p) //to print representation of a pointer.
	fmt.Printf("|%6d|%6d|\n",12,345)
	fmt.Printf("|%-6.2f|%-6.2f|\n",1.2,3.45) //to left-justify use -flag

	s:= fmt.Sprintf("a %s","string") // Sprintf formats and return a string 
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s \n","error") // format + print to io.Writers other than os.Stdout using Fprintf
}