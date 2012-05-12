// based on http://golang.org/pkg/regexp/#Regexp.FindAllString
package main
 
import (
  "fmt"
   "regexp"
 
)
func main(){

  r,_:=regexp.Compile("p|oo")
 
 matches:=r.FindAllString("oop",-1)
 // %q quote element to see... instead of %s
 fmt.Printf("%q \n",matches)
 
 for _,m:=range matches {
  fmt.Println(m)
 }

}