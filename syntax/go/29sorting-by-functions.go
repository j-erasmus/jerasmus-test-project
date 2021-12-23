/*
we can sort a collection by something other than its natural order like length

To sort by custom function in Go, we need corresponding type.
we need to implement sort.Interface - Len , Less, Swap methods
*/

package main 

import (
	"fmt"
	"sort"
)

type ByLength []string

func(s ByLength) Len()int{
	return len(s)
}

func(s ByLength) Swap(i,j int){
	s[i],s[j] = s[j],s[i]
}

func(s ByLength) Less(i,j int)bool {
	return len(s[i])< len(s[j])
}

func main(){
	fruits := []string{"peach","banana","kiwi","watermalon"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}