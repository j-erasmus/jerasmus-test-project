package main 

import "fmt"
import "net"
import "net/url"

func main(){

	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	//s:= "www.google.com" //gives error
	u,err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)  // Accessing the scheme

	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p,_ := u.User.Password()
	fmt.Println(p)

	fmt.Println(u.Host) // Host contains hostname and the port if present
	host,port,_ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Path) //extract path and fragment after #
	fmt.Println(u.Fragment)

	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])

}