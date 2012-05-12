// call ruby from go
package main

import (
    "fmt"
    "os/exec"
)

func main() {
 out,_:=exec.Command("ruby","-e","p rand").Output()
 fmt.Println(string(out))
}