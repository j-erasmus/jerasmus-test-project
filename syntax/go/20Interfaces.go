//interfaces are named collection of method signatures.

package main 

import "fmt"
import "math"

type geometry interface {
	area() float64
	perim() float64
}

type rect struct{
	width,height float64
}

type circle struct{
	radius float64
}

/*to implement geometry interface we need to implement its methods
   i.e. write methods with receiver as rect and circle.*/

func (r rect)area() float64{
	return r.width * r.height
}

func(r rect)perim() float64{
	return 2*r.width + 2*r.height
}

func (c circle)area() float64{
	return math.Pi * math.Pow(c.radius,2)
}

func (c circle)perim() float64{
	return 2*math.Pi*c.radius
}

func measure(g geometry){
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main(){
	var r geometry
	r = rect{20,30}

	var c geometry
	c = circle{5}

	fmt.Printf("%T %T \n",r,c)
	measure(r)
	measure(c)
}