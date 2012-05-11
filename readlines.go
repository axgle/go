/**
golang: read text file into string array (and write)

"The ability to read (and write) a text file into and out of a string array is  a fairly common requirement
http://stackoverflow.com/questions/5884154/golang-read-text-file-into-string-array-and-write
"
here is a simple that quick coding . 
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