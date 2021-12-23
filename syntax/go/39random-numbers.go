package main 

import "time"
import "fmt"
import "math/rand"

func main(){

	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	fmt.Print((rand.Float64()*5)+5,",") // rand floats in 5.0 to 10.0

	s1 := rand.NewSource(time.Now().UnixNano()) // default no generator is deterministic ,so it'll produce same no each time
	r1 := rand.New(s1)                    //so for varying sequences give it a seed that changes.
										  // you can use crypto/rand 

	fmt.Print(r1.Intn(100),",")
	fmt.Print(r1.Intn(100))
	fmt.Println()
}