package main

import (
	"fmt"
	"syscall"
)

type (
	HANDLE        uintptr
	DWORD         uint32
	HWND          HANDLE
	HWINEVENTHOOK HANDLE
	LONG          int32
)

var (
	moduser32           = syscall.NewLazyDLL("user32.dll")
	procSetWinEventHook = moduser32.NewProc("SetWinEventHook")
)

const EVENT_SYSTEM_FOREGROUND = 0x0003

type WINEVENTPROC func(HWINEVENTHOOK, DWORD, HWND, LONG, LONG, DWORD, DWORD)

func windowLogger(hook HWINEVENTHOOK, event DWORD, hwnd HWND, idObject LONG, idChild LONG, dwEventThread DWORD, dwmsEventTime DWORD) {
	if event == EVENT_SYSTEM_FOREGROUND {
		fmt.Println("we passed")
	}
}

func main() {
	fmt.Println("Starting")
	_ = syscall.NewCallback(windowLogger) // this will panic with message: "compileCallback: expected function with one uintptr-sized result"
}
