package main

import (
    "fmt"
	"syscall"
	"unsafe"
 
	"github.com/AllenDang/w32"
)

func StringToUTF16PtrLen(s string) (ptr uintptr, l uint16) {
	u := syscall.StringToUTF16(s)
	l = uint16(len(u) * 2) //size in uint8, length of uint16
	ptr = uintptr(unsafe.Pointer(&u[0]))
	return
}

func RegSetKey(hKey w32.HKEY, subKey string, valueName string,data interface{}) (errno int) { 

	var dataPtr  uintptr
	var dataType w32.DWORD
	var dataLen uint16
	
	switch v := data.(type) {
	case uint32:
		dataPtr, dataLen = uintptr(unsafe.Pointer(&v)), 4
		dataType = w32.REG_DWORD
	case string:
		// The comment on MSDN regarding escaping back-slashes, are c-lang specific.
		// The API just takes a normal NUL terminated string, no special escaping required.
		dataPtr, dataLen = StringToUTF16PtrLen(v)
		dataType = w32.REG_SZ
	}

    return w32.RegSetKeyValue(hKey, subKey , valueName , dataType, dataPtr, dataLen)
}


func main() {
 fmt.Println(w32.RegGetString(w32.HKEY_CURRENT_USER,"environment","path")) 
 RegSetKey(w32.HKEY_CURRENT_USER,"environment","path_test_say","hello world")
 fmt.Println(w32.RegGetString(w32.HKEY_CURRENT_USER,"environment","path_test_say")) 
}