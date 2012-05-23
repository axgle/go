package main

import "fmt"
import "crypto/md5"
func main() {
   s:=md5.New()
   s.Write([]byte("test"))
	fmt.Printf("%x",  s.Sum(nil) )
}