
package main
/*
#include <Windows.h>
#include "AutoIt3.h"
#cgo LDFLAGS:./AutoItX3.dll
void run(){
  AU3_Run(L"notepad.exe",L"",1);
  AU3_Sleep(500);
  AU3_WinActivate(L"无标题 - 记事本",L"");
  AU3_Send(L"中文{enter}",0);
  AU3_ClipPut(L"试试剪贴板");
}
*/
import "C"
import(
 //"fmt"
 //"syscall"
)
func main(){
  C.run()
 
 //fmt.Printf("hi")

}
 