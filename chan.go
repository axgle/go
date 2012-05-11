package main
import "fmt"
import "time"
func push(ch chan int){
 for i:=0;;i++{
   ch<-i
 }
}

func pop(ch chan int){
 for{
   fmt.Println(<-ch)
 }

}
 func main(){
 
  ch:=make(chan int)
  go push(ch)
  go pop(ch)
  time.Sleep(3*time.Second)

}