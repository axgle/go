package main
import "fmt"
import "os/exec"
import "encoding/json"
func main(){


 b,_:=exec.Command("php","-r",`echo json_encode(glob("*"));`).Output()

 fmt.Printf("%s\n",b)
  
  var data []string
  json.Unmarshal(b,&data)
 
 
  fmt.Printf("%+q",data)
}