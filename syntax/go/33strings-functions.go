/*
strings package provides many useful string related functions.
*/
package main 

import s "strings"
import "fmt"

var p = fmt.Println

func main(){
	p("contains:   ", s.Contains("test","es"))
	p("Count:      ", s.Count("test","t"))
	p("Has Prefix: ", s.HasPrefix("test","te"))
	p("Has Suffix: ",s.HasSuffix("test","st"))
	p("Index:      ", s.Index("test","e"))
	p("Join:       ",s.Join([]string{"a","b"},"-"))
	p("repeat:     ",s.Repeat("a",5))
	p("replace:    ",s.Replace("foo","o","0",-1))
	p("Split:      ",s.Split("a-b-c-d-e","-"))
	p("ToLower:    ",s.ToLower("TEST"))
	p("ToUpper:    ",s.ToUpper("test"))
	p()
	p("Len:        ",len("hello"))
	p("Char:       ","hello"[1])
}