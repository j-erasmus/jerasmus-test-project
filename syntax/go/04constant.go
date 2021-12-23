package main 

import ( 
	"fmt"
	"math"
)

const s string = "constant"  //declaring constant

func main(){ 
	fmt.Println(s)

	const b = 500000000   //const can appear at place of var

	const c = 3e20/b
	fmt.Println(c)

	fmt.Println(int64(c))   //const has no type untill its given

	fmt.Println(math.Sin(c))  //here math.Sin expects float64

}