package main

import "syscall"

// http://www.dreamincode.net/forums/topic/148684-turn-off-your-lcd/
func main() {
	u := syscall.MustLoadDLL("user32.dll")
	win, _, _ := u.MustFindProc("FindWindowA").Call(0, 0)
	u.MustFindProc("SendMessageA").Call(win, 0x112, 0xf170, 2)

}
