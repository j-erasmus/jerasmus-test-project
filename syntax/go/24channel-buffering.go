/*
by default channels are unbuffered, 
i.e. they will only accept sends(chan <-) if there is corresponding receive(<-chan)ready to recieve sent value.
Buffered channels accept a limited number of values without corresponding reciever 
*/
// not understood exactly 
package main 

import "fmt"

func main(){
	message := make(chan string,2)

	message <- "buffered"
	message <- "chaneel"

	fmt.Println(<-message)
	fmt.Println(<-message)


}