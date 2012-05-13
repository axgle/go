// based on http://btl-com.net/gcc-4.7.0/libgo/go/mime/type_windows.go
// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"syscall"
	"unsafe"
	"fmt"
)

func main() {
	var root syscall.Handle
	
	error:=syscall.RegOpenKeyEx(syscall.HKEY_CURRENT_USER, syscall.StringToUTF16Ptr(`Control Panel\Accessibility\HighContrast`),
								0, syscall.KEY_READ, &root)
		
	defer syscall.RegCloseKey(root)
	
	var buf [1 << 10]uint16
	
	var typ uint32
	n := uint32(len(buf) ) // api expects array of bytes, not uint16
	if syscall.RegQueryValueEx(
		root, syscall.StringToUTF16Ptr(`Previous High Contrast Scheme MUI Value`),
		nil, &typ, (*byte)(unsafe.Pointer(&buf[0])), &n) != nil {		 
		 
	}
	
	path:=syscall.UTF16ToString(buf[:])
	
	fmt.Println(path)		
	
	if error != nil {
		fmt.Println(error)
			return
	}
	
	 
}


func getPath() {
	var root syscall.Handle
	
	error:=syscall.RegOpenKeyEx(syscall.HKEY_CURRENT_USER, syscall.StringToUTF16Ptr(`Environment`),
								0, syscall.KEY_READ, &root)
		
	defer syscall.RegCloseKey(root)
	
	var buf [1 << 10]uint16
	
	var typ uint32
	n := uint32(len(buf) ) // api expects array of bytes, not uint16
	if syscall.RegQueryValueEx(
		root, syscall.StringToUTF16Ptr(`Path`),
		nil, &typ, (*byte)(unsafe.Pointer(&buf[0])), &n) != nil {		 
		 
	}
	
	path:=syscall.UTF16ToString(buf[:])
	
	fmt.Println(path)		
	
	if error != nil {
		fmt.Println(error)
			return
	}
	
	 
}

 