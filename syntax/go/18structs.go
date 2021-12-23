/*
	structs are typed collections of fields.
	Useful for grouping data together to form records.

*/

package main 

import "fmt"

type person struct {
	name string
	age int
}

func main(){

	fmt.Println(person{"tushar",22})
	fmt.Println(person{name:"reshma",age:27})

	s := person{"hemant",21}
	fmt.Println(s)
	fmt.Println(&s)

	s.age = 22
	fmt.Println(s)
}