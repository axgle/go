// http://msdn.microsoft.com/en-us/library/ms683169(VS.85).aspx
package main
import(
 "fmt"
)
//#include <Windows.h>
import "C"
func main(){

 //c:\>chcp
 code:=C.GetConsoleOutputCP()
 fmt.Printf("%v\n",code)
 
 //c:\>chcp 65001
 //65001 for utf8
 C.SetConsoleOutputCP(65001)//default 936 for gbk

 new_code:=C.GetConsoleOutputCP()
 fmt.Printf("%v",new_code)
 
 C.SetConsoleOutputCP(code)
 
}
 