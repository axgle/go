package main
import "strings"
import "os/exec"
func main(){
  cmd:= strings.Fields(`wscript demo.vbs`)
  cmd=strings.Fields(`reg query HKCU\environment`)
  s,_:=exec.Command(cmd[0],cmd[1:]...).Output()
  println(string(s))
}