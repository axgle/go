
package main
/*
#include <Windows.h>
#include "AutoIt3.h"
#cgo LDFLAGS:./AutoItX3.dll
void run(){
  AU3_Run(L"notepad.exe",L"",1);
}
*/
import "C"
import(
 "fmt"
 //"syscall"
)
func main(){
  C.run()
 
 fmt.Printf("hi")

}
 