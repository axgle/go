// http://answers.yahoo.com/question/index?qid=20070303145917AANhnte
package main
import(
 //"fmt"
)
/*
#include <Windows.h>
#include <stdio.h>
void run(){
	TCHAR currentTitle[512];
	GetConsoleTitle(currentTitle, sizeof(currentTitle) / sizeof(TCHAR));
	printf(currentTitle);
}
*/
import "C"
func main(){
 C.run() 
}
 