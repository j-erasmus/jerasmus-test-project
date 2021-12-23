package main 

import "fmt"

func main(){ 
	nums := []int{2,3,4}
	sum :=0
	for _,num := range nums{ 
		sum+=num
	}
	fmt.Println("sum:",sum)

	for i,num := range nums{ 
		fmt.Println("index:",i," num:",num)
	}

	kvs := make(map[string]string)
	kvs["a"]= "apple"
	kvs["b"]= "banana"

	for k,v := range kvs{
		fmt.Println(k," -> ",v)
	}

	for i,c := range "ao"{
		fmt.Println(i,c)
	}
}