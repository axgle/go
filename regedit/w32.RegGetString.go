package main

import (
    "fmt"
	"github.com/AllenDang/w32"
)

func main() {
 fmt.Println(w32.RegGetString(w32.HKEY_CURRENT_USER,"environment","path")) 
}