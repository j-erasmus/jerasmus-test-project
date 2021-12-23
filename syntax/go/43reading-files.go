package main 

import(
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error){ //for checking most of errors
	if e != nil{
		panic(e)
	}
}

func main(){

	dat, err := ioutil.ReadFile("/home/tushar/Desktop/tushar/go_work/go_by_examples/test.txt")    //reading file
	check(err)
	fmt.Print(string(dat))

	f, err := os.Open("/home/tushar/Desktop/tushar/go_work/go_by_examples/test.txt")
	check(err)

	b1 := make([]byte,5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s \n", n1, string(b1))

	o2, err := f.Seek(6,0)
	check(err)
	b2 := make([]byte,2)
	n2,err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s \n",n2,o2,string(b2))

	o3,err := f.Seek(6,0)
	check(err)
	b3 := make([]byte,2)
	n3,err := io.ReadAtLeast(f,b3,2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s \n",n3,o3,string(b3))

	_, errr := f.Seek(0,0)
	check(errr)

	r4 := bufio.NewReader(f)  // fufio implements buffered reader that may be useful for efficiency with small reads and
	b4, err := r4.Peek(5)     // it also provides addtional reading methods.
	check(err)
	fmt.Printf("5 bytes: %s \n", string(b4))

	f.Close()     // close the file usually this would be scheduled immediately after opening with defer
}

