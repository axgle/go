package main

/*
#include "windows.h"
void run(){

	ShowWindow(FindWindowA("Shell_TrayWnd",""),0);//0 for hide, 5 for show
	Beep(2500,3000);//3000 for 3 seconds
	ShowWindow(FindWindowA("Shell_TrayWnd",""),5);
}
*/
import "C"

func main() {

	C.run()

}
