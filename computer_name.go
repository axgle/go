package main
import "syscall"

func main() {
  name,_:=syscall.ComputerName()
  println(name)
}