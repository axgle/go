package main
import(
"os"
"os/exec"
"strings"
"time" 
)

func main(){  
  if len(os.Args)>1 {
  // >=php5.4 buildin web server
  cmd := "php -S localhost:"+os.Args[1] 
  cmd_fields:=strings.Fields(cmd)
  go func(){
   exec.Command(cmd_fields[0],cmd_fields[1:]...).Run()
  }()
 
  time.Sleep(5*time.Second)
  
  } else{
   println("pls input port")
  }
}