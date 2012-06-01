package main
import(
 "fmt"
 "syscall"
 "unsafe"
)
func PathFileExists(p string) bool{
 shlwapi:=syscall.MustLoadDLL("shlwapi.dll")
 PathFileExistsProc:=shlwapi.MustFindProc("PathFileExistsW")
 r,_,_:=PathFileExistsProc.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(p))))
 return r>0
}
func main(){ 
 fmt.Printf("%v",PathFileExists("c:\\"))
 fmt.Printf("%v",PathFileExists("PathFileExists.go"))

}
 