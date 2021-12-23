package main 
 
import (
	"fmt"
	"sort"
)

func main(){

	strs := []string{"c","a","b"}
	sort.Strings(strs)
	fmt.Println("strings:",strs)

	ints := []int{4,7,2,1,3,4}
	sort.Ints(ints)
	fmt.Println("ints:",ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("sorted:",s)
}