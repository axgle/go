package main
import ("fmt"
"os"
"bufio"
)


func main() {
    
	f,_:=os.Open("main.go")
	r:=bufio.NewReader(f)
	defer f.Close()
	
	line,_:=r.ReadString('\n')
	line,_=r.ReadString('\n') 
	line,_=r.ReadString('\n') 
    fmt.Printf(line)
	
   if len(os.Args)>1{
      fmt.Println(os.Args[1])
   }
 
 }
 // http://stackoverflow.com/questions/6141604/go-readline-string