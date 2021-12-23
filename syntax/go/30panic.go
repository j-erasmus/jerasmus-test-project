/*
panic means something went unexpectedly wrong.
we use it to fail fat on errors that shouldn't occur during normal operation or we aren't prepared for.
*/

package main 

import "os"

func main(){
	panic("a problem")

	_,err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}