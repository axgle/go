package main
import(
 "os"
)
func main(){
 f,e:=os.Create("demo.go")
 if e!=nil{
 }
 
 f.WriteString(
`package main
import(
 "fmt"
)
func main(){

 fmt.Printf("hi")

}
 `)

}