package main 

import "strconv"
import "fmt"

func main(){

	f,_ := strconv.ParseFloat("1.2345",64)
	fmt.Println(f)

	i,_ := strconv.ParseInt("123",0,64) // 0 means infer base from string, 64- fit in 64 bits 
	fmt.Println(i)

	d,_ := strconv.ParseInt("0x1c8",0,64)
	fmt.Println(d)

	u,_ :=strconv.ParseUint("789",0,64)
	fmt.Println(u)

	k,_ := strconv.Atoi("135") // aAtoi for basic base 10 int parsing
	fmt.Println(k)

	_,e := strconv.Atoi("wat")
	fmt.Println(e)
}