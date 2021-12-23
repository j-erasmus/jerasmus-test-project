/*
Channels are the pipes that connect concurrent goroutinges.
You can send values into channels from one goroutine and
receive those values into another goroutine.
*/

package main 

import "fmt"

func main(){

	message := make(chan string)  // create a channel with make(chan val-type)

	
	go func(){ 
		message <- "ping"   //send a value to channel with channel <- syntax
	}()

	msg := <-message    //receives value from channel
	fmt.Println(msg)
}
