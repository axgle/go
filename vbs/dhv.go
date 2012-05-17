package main
import "strings"
import "os/exec"
func main(){
  cmd:= strings.Fields(`wscript fso.vbs`)
  exec.Command(cmd[0],cmd[1:]...).Run()
}