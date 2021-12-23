/*
we need our programs to perform operations of collection of data
Go does not support generics;
in Go its common to provide collection functions if and when they 
are needed for prg and data types.

some eg collection functions for slices of strings.

need to see again
*/

package main 

import (
	"strings"
	"fmt"
)

func Index(vs []string, t string) int{ //return first index of target index or -1 for not found
	for i,v := range vs{
		if v == t{
			return i
		}
	}
	return -1
}

func Include(vs []string, t string) bool{ // returns true if the target string t is in the slice
	return Index(vs,t) >=0
}

func Any(vs []string, f func(string) bool) bool{ // return true if one of the strings satisfies the predicate f.
	for _,v := range vs{
		if f(v){
			return true
		}
	}
	return false
}

func All(vs []string, f func(string)bool) bool{ // returns true if all strngs satisfied the predicate f.
	for _,v := range vs{
		if !f(v){
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool)[]string{ //return a new slice containing result of applying fn f to each stirng
	vsf := make([]string,0)
	for _,v := range vs{
		if f(v){
			vsf = append(vsf,v)	
		}
	}
	return vsf
}

func Map(vs []string, f func(string) string) []string{
	vsm := make([]string, len(vs))
	for i,v := range vs{
		vsm[i]= f(v)
	}
	return vsm
}
func main(){
	var strs =[]string{"peach","apple","pear","plum"}

	fmt.Println(Index(strs,"pear"))

	fmt.Println(Include(strs,"grape"))

	fmt.Println(Any(strs, func(v string) bool{
			return strings.HasPrefix(v,"p")
		}))

	fmt.Println(All(strs,func(v string) bool{
			return strings.HasPrefix(v,"p")
		}))

	fmt.Println(Filter(strs, func(v string) bool{
			return strings.Contains(v,"e")
		}))

	fmt.Println(Map(strs, strings.ToUpper))
}

