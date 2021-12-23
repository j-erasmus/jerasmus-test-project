/*
sometimes we like our go prgs to handle unix signals.
*/
package main 

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main(){
	sigs := make(chan os.Signal,1)
	done := make(chan bool,1)
	//Go signal notification works by sending os.Signal values
	//we'll create a channel to receive these notifications 

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	//signal.Notify registers given channel to receive notifications of specified signals.

	go func(){          //this goroutine executes a blocking recive for signals.
		sig := <-sigs   // when it gets one it'll print it out and then notify prg that it can finish.
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}