package main
 
/*
#include <malloc.h>
#include <string.h>
const char* file(){
 char* f;
 f=malloc(strlen(__FILE__)+1);
 strcpy(f,__FILE__);
 return f;
}
*/
import "C"
import "unsafe"
func main() {
   cs:=C.file() 
   
   defer C.free(unsafe.Pointer(cs))
   
   println( C.GoString(cs))
   
 
 
}

 


