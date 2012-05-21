// display gbk/gb2312 Chinese chars on window's cmd which default using gbk encoding
package main
import "fmt"
import "code.google.com/p/mahonia"
func main(){
  enc:=mahonia.NewEncoder("gbk")
  fmt.Println(enc.ConvertString("hello,世界"))  
}