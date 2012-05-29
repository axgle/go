package main
import "fmt"

// extern char binary_hello_txt_start[];
import "C"

func main(){
  fmt.Println("hi"+C.CString(C.binary_hello_txt_start))
}