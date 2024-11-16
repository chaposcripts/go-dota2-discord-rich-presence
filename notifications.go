package main

import (
	"gopkg.in/toast.v1"
	"syscall"
	"unsafe"
)

type MessageBoxType int

const (
	MessageBoxTypeAbortRetryIgnore  MessageBoxType = 0x00000002
	MessageBoxTypeCancelTryContinue MessageBoxType = 0x00000006
	MessageBoxTypeHelp              MessageBoxType = 0x00004000
	MessageBoxTypeOk                MessageBoxType = 0x00000000
	MessageBoxTypeOkCancel          MessageBoxType = 0x00000001
	MessageBoxTypeRetryCancel       MessageBoxType = 0x00000005
	MessageBoxTypeYesNo             MessageBoxType = 0x00000004
	MessageBoxTypeYesNoCancel       MessageBoxType = 0x00000003
)

func ShowMessageBox(title string, text string, style MessageBoxType) {
	if len(title) > 0 {
		title = " / " + title
	}
	syscall.NewLazyDLL("user32.dll").NewProc("MessageBoxW").Call(0, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))), uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Dota2 Discord Rich Presence"+title))), uintptr(style))
}

func ShowToast(title string, message string) error {
	notification := toast.Notification{
		AppID:   "Dota2 Discord Rich Presence",
		Title:   title,
		Message: message,
	}
	return notification.Push()
}
