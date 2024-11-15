package main

import "gopkg.in/toast.v1"

func ShowToast(title string, message string) error {
	notification := toast.Notification{
		AppID:   "Dota2 Discord Rich Presence",
		Title:   title,
		Message: message,
	}
	return notification.Push()
}
