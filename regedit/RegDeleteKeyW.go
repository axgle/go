package main
import "fmt"
 import "syscall"
import "unsafe"
const HKEY_CURRENT_USER = 0x80000001

var f=fmt.Printf


 // C:\>reg add HKCU\google_test
 //C:\>reg delete HKCU\google_test
func main() {
  var a = syscall.MustLoadDLL("advapi32")
  var regdeletekeyw = a.MustFindProc("RegDeleteKeyW")
  
  _, _, e1 :=regdeletekeyw.Call(
   uintptr(HKEY_CURRENT_USER),
   uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(`google_test`))),
  )
  f("%v",e1)
  
}
//more see https://github.com/axgle/service/blob/master/service_windows.go
//https://github.com/AllenDang/w32/blob/master/constants.go