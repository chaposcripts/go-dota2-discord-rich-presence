package main

import (
	"syscall"
	"unsafe"
)

func ShowMessageBox(title string, text string) {
	user32 := syscall.NewLazyDLL("user32.dll")
	messageBox := user32.NewProc("MessageBoxW")
	const (
		MB_OK              = 0x00000000
		MB_ICONINFORMATION = 0x00000040
	)
	messageBox.Call(0, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))), uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))), MB_OK|MB_ICONINFORMATION)
}
