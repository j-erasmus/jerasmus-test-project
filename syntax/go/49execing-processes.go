/*
In spawning external processes - we need an external process accessible to running Go process.
But somethimes we watn to completely replace current Go process with another one.
we'll use exec function
*/

package main 

import "syscall"
import "os"
import "os/exec"

func main(){

	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil{
		panic(lookErr)
	}

	args := []string{"ls", "-a","-l","-h"}

	env := os.Environ()

	execErr := syscall.Exec(binary,args,env)
	if execErr != nil{
		panic(execErr)
	}
}