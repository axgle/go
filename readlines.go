package main

import(
 "fmt"
 "strings"
 "io/ioutil"
)

func main(){  
   d,_:=ioutil.ReadFile("readlines.go")
   lines:=strings.Split(string(d),"\n")
   for i,l:= range lines{   
	fmt.Printf("%d %s",i,l)
   }
   
}