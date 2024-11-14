package main

import (
	"fmt"
	"os"
	"time"

	"github.com/getlantern/systray"
)

const (
	port                 string = "3388"
	discordApplicationID string = "751932819676332042"
	baseImageUrl         string = "https://www.dotabuff.com/assets/heroes/%s.jpg" //"https://cdn.akamai.steamstatic.com/apps/dota2/images/dota_react/heroes"
	errorImageUrl        string = ""
)

var activity DotaGsiRequest
var checkbox *systray.MenuItem

func onReady() {
	systray.SetTitle("Dota2 Discord Rich Presence")
	systray.SetTooltip("Dota2 Discord Rich Presence\nt.me/chaposcripts")
	checkbox = systray.AddMenuItemCheckbox("Enabled", "Toggle status", true)
	go func() {
		for {
			<-checkbox.ClickedCh
			if checkbox.Checked() {
				checkbox.Uncheck()
				LogOut()
			} else {
				checkbox.Check()
				LogIn(discordApplicationID)
			}
		}
	}()
	mQuitOrig := systray.AddMenuItem("Quit", "Quit")
	go func() {
		<-mQuitOrig.ClickedCh
		fmt.Println("Closing the application...")
		systray.Quit()
		os.Exit(0)
	}()
}

func main() {
	go systray.Run(onReady, func() {})
	fmt.Println("Started, connecting to discord...")

	LogIn(discordApplicationID)

	HandleGSI(port)
	for {
		if checkbox.Checked() {
			fmt.Println("Updating RP...")
			Update()
		} else {
			fmt.Println("Skipping update (disabled)")
		}
		time.Sleep(5 * time.Second)
	}
}
