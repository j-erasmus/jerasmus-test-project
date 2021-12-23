package main 

import "fmt"

func ping(pings chan<- string, msg string){  //ping func accepts a channel for sending values.
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string){ //pong func accepts one channel for receives and second for sends
	msg := <-pings
	pongs <- msg
}

func main(){
	pings := make(chan string, 1)
	pongs := make(chan string,1)
	ping(pings,"passed message")
	pong(pings,pongs)
	fmt.Println(<-pongs)
}