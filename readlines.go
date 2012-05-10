/**
golang: read text file into string array (and write)

"The ability to read (and write) a text file into and out of a string array is  a fairly common requirement
http://stackoverflow.com/questions/5884154/golang-read-text-file-into-string-array-and-write
"
here is a simple that quick coding . write less do more,dont argument performance,it just small task
just do it and done quickly like Ruby!

why does  io.Readlines NOT in the golang Default? why I must write more codes do it just for performance?
in most cases,who care? Wasting developer's time is not good.
I hope go lang can be both have performance ,and easy quickly use like ruby
*/

package main

import(
  "io/ioutil"
 "strings"
 "fmt"
)

func main(){  
   data,_:=ioutil.ReadFile("readlines.go")
   lines:=strings.Split(string(data),"\n")
   for i,l:= range lines{   
	fmt.Printf("%d %s",i,l)
   }
   
}