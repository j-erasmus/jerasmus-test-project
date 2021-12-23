package main 

import "fmt"

type rect struct{
	width,height int
}

func (r *rect) area() int{     // here receiver type is pointer to rect so we could change value of rect and
	//r.width =10              // it sustain value throutout program.
	return r.width*r.height
}

func (r rect) perim() int{         // receiver type is just rect so we could not change values of rect permanantely.

	return 2*r.width + 2*r.height
}

func main(){

	r:= rect{20,10}

	fmt.Println("area:",r.area())
	fmt.Println("perimeter:",r.perim())
}