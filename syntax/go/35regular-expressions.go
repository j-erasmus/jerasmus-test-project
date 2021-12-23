/*
Go offers support for regular expressions.
*/

package main 

import(
	"bytes"
	"fmt"
	"regexp"
)

func main(){

	match,_ := regexp.MatchString("p([a-z]+)ch","peach")
	fmt.Println(match)

    r,_ := regexp.Compile("p([a-z]+)ch")

    fmt.Println(r.MatchString("peach"))

    fmt.Println(r.FindString("peach punch")) // finds the match for regexp.

    fmt.Println(r.FindStringSubmatch("peach punch")) //includes info about whole pattern matches and submatches within those matches
                                     // eg.this will return info for p([a-z]+)ch and ([a-z]+)
    //r.FindStringSubmatchIndex

    fmt.Println(r.FindAllString("peach punch pinch",-1)) // applies to all matches in the input not just first
    //r.FindAllStringSubmatchIndex

    fmt.Println(r.FindAllString("peach punch pinch",2)) // non-negative integer will limit no of matches.

    fmt.Println(r.Match([]byte("peach"))) //we can provide byte[] - drop string from function name.

    r = regexp.MustCompile("p([a-z]+)ch") //MustCompile for creating constant with reg exp.
    fmt.Println(r)                        //plain compile won't work because it has two return values.

    fmt.Println(r.ReplaceAllString("a peach","<fruit>")) // used to replace subsets of strings with other values.
	
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in,bytes.ToUpper)
	fmt.Println(string(out))    

}
