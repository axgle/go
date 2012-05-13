package main
import "syscall"
func main(){

  dhv:=syscall.MustLoadDLL("dhv.dll")
  hi:=dhv.MustFindProc("hi")
  s,_,_:=hi.Call()
  println(s)

}