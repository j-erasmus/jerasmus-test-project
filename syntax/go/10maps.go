package main 

import "fmt"

func main(){ 

	m:= make(map[string]int)

	m["a"] =1
	m["b"] =2

	fmt.Println("map:",m)

	fmt.Println("len",len(m))

	delete(m,"a")
	fmt.Println("map:",m)

	_,prs := m["a"]
	if prs {
		fmt.Println("present")
	}else {

		fmt.Println("not present")
	}
	fmt.Println("prs:",prs)
}