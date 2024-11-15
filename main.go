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
)

var checkbox *systray.MenuItem
var lastUpdate time.Time

func onReady() {
	systray.SetTemplateIcon(logo, logo)
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
				LogIn(discordApplicationID, false)
			}
		}
	}()
	quitButton := systray.AddMenuItem("Quit", "Quit")
	go func() {
		<-quitButton.ClickedCh
		fmt.Println("Closing the application...")
		systray.Quit()
		os.Exit(0)
	}()
}

func main() {
	go systray.Run(onReady, func() {})
	fmt.Println("Started, connecting to discord...")
	LogIn(discordApplicationID, true)
	go HandleGSI(port)
	for {
		time.Sleep(5 * time.Second)
		if loggedIn && time.Now().Unix()-lastUpdate.Unix() > 20 {
			LogOut()
			fmt.Println("Logged out (no new updates in 20 seconds)")
			ShowToast("Rich Presence disabled", "No game information received in past 20 seconds. Rich presence was automatically disabled!")
		}
	}
}
