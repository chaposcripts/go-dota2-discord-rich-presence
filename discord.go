package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/hugolgst/rich-go/client"
)

var loggedIn bool = false

func LogIn(clientId string, showNotification bool) error {
	if loggedIn {
		return errors.New("ALREADY_LOGGED_IN")
	}
	err := client.Login(clientId)
	if err == nil {
		loggedIn = true
		if showNotification {
			ShowMessageBox("Dota2 Discord Rich Presence", "Connected to discord!\nYou can close/disable app in system tray!\nt.me/chaposcripts", MessageBoxTypeOk)
		}
	}
	return err
}

func LogOut() error {
	if !loggedIn {
		return errors.New("NOT_LOGGED_IN")
	}
	client.Logout()
	loggedIn = false
	return nil
}

func Update(activity DotaGsiRequest) {
	if !loggedIn {
		LogIn(discordApplicationID, false)
		fmt.Println("Logging in...")
		return
	}

	timestamp := time.Unix(time.Now().Unix()-int64(activity.Map.ClockTime), 0) //time.Unix(time.Now().Unix()-int64(activity.Map.ClockTime), 0)
	// fmt.Println("TIMESTAMP", timestamp.Unix(), "MapClock", activity.Map.ClockTime, "NOW", time.Now().Unix())
	var settings client.Activity
	if len(activity.Map.GameState) > 0 {
		var isInGame = activity.Map.GameState == string(DotaGameStateTeamShowcase) || activity.Map.GameState == string(DotaGameStateGameInProgress) || activity.Map.GameState == string(DotaGameStatePreGame)
		if isInGame {
			settings = client.Activity{
				State:      fmt.Sprintf("%d / %d / %d", activity.Player.Kills, activity.Player.Deaths, activity.Player.Assists),
				Details:    fmt.Sprintf("%s (LVL: %d)", fixName(activity.Hero.Name), activity.Hero.Level),
				LargeImage: getHeroImageUrl(activity.Hero.Name),
				LargeText:  getItemsAsString(activity.Items),
				SmallImage: ImageDotaLogo,
				SmallText:  fmt.Sprintf("Radiant/Dire: %d/%d\nMatch ID: %s\nPlaying for %s", activity.Map.RadiantScore, activity.Map.DireScore, activity.Map.MatchID, activity.Player.TeamName),
				Timestamps: &client.Timestamps{
					Start: &timestamp,
				},
			}
		} else {
			settings = client.Activity{
				State:      DotaGameStateLabel[DotaGameState(activity.Map.GameState)],
				LargeImage: DotaGameStateImage[DotaGameState(activity.Map.GameState)],
				SmallImage: ImageDotaLogo,
			}
		}
	} else {
		settings = client.Activity{
			Details:    "In Main Menu",
			LargeImage: ImageDotaLogo,
		}
	}
	client.SetActivity(settings)
	lastUpdate = time.Now()
}
