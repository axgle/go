package main
import(
"fmt"
"os/exec"
"strings"
)

func main(){

 cmd:=`reg query HKCU\environment -f path`
 //cmd = "ruby -e 'p rand'"
 
 //https://groups.google.com/group/golang-nuts/browse_thread/thread/af9d65c96669364c
 argv:=strings.Fields(cmd)
 d,_:=exec.Command(argv[0],argv[1:]...).Output()
 
 fmt.Println(string(d))
}