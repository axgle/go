// call c from go lang 
// http://code.google.com/p/go-wiki/wiki/cgo
package main
 
 /*
#include <stdio.h>
#include <malloc.h>
*/
import "C"
import "unsafe"
 
 func main(){
  
  cs:=C.CString(`
  wokao
  `)
  C.puts((cs));
  
  defer C.free(unsafe.Pointer(cs))
  
 
}