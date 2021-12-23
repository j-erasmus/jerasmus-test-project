package main 

import "fmt"

func main(){ 

	i:=1
	for i<3 { 
		fmt.Println(i)
		i++
	}

	for j:=4;j<7;j++{ 
		fmt.Println(j)
	}

	for { 
		fmt.Println("true loop")
		break
	}

	for n:=0;n<=5;n++{ 
		if n%2 ==0{ 
			continue
		}
		fmt.Println(n)
	}
}