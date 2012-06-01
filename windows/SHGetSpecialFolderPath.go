package main

/*
#include <shlobj.h> 
#include <stdio.h>

void run(){
    TCHAR szPath[MAX_PATH];  
    SHGetSpecialFolderPath(NULL, szPath, CSIDL_RECENT, 0);  
    printf(szPath);  
    SHGetSpecialFolderPath(NULL, szPath, CSIDL_DESKTOP, 0);  
    printf(szPath);  	
}
*/
import("C")
 
func main(){  
 C.run()
}
 