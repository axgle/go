package main

import (
    "fmt"
    "os"
)

func main() {
    f,_:=os.Open("filesize.go")
	s,_:=f.Stat()
	fmt.Println(s.Size())
	f.Close() 
}