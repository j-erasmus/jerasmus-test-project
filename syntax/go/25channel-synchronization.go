package main 

import "fmt"
import "time"

//this fn we'll run in a goroutine. done channel will be used to notify 
//another goroutine tha this function's work is done
func worker(done chan bool){  
	fmt.Print("working ...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true

}

func main(){
	done := make(chan bool,1)
	go worker(done)

	<-done
}