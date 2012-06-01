package main
import(
 "fmt"
 "syscall"
)
const(
 SHERB_NOCONFIRMATION = 0x00000001
 SHERB_NOPROGRESSUI = 0x00000002
 SHERB_NOSOUND = 0x00000004
)
func main(){
 d:=syscall.MustLoadDLL("shell32.dll")
 SHEmptyRecycleBinProc:=d.MustFindProc("SHEmptyRecycleBinW")
 SHEmptyRecycleBinProc.Call(uintptr(0),uintptr(0),SHERB_NOCONFIRMATION)
 fmt.Printf("cleaned")

}
 