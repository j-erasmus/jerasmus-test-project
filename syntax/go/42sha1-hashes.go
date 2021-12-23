/*
SHA1 hashes are used to compute short identities for binary or text blobs.
It identify versioned files and directories.
*/

package main 

import "crypto/sha1"
import "fmt"

func main(){

	s := "sha1 this string"
	h := sha1.New()

	h.Write([]byte(s)) // write expects bytes 

	bs := h.Sum(nil) // finalized hash result as byte slice.

	fmt.Println(s)

	fmt.Printf("%x\n",bs)


}