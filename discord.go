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
	} else {
		if showNotification {
			ShowMessageBox("Dota2 Discord Rich Presence // Error", "Error connecting to Discord:\n"+err.Error(), MessageBoxTypeOk)
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

	var timestamp time.Time = time.Unix(time.Now().Unix()-int64(activity.Map.ClockTime), 0)
	var settings client.Activity = client.Activity{}
	if len(activity.Map.GameState) > 0 {
		settings.SmallImage = ImageDotaLogo
		// var isInGame = activity.Map.GameState == string(DotaGameStateTeamShowcase) || activity.Map.GameState == string(DotaGameStateGameInProgress)
		switch activity.Map.GameState {
		case string(DotaGameStatePostGame):
			var endGameText string = "Loose"
			if activity.Player.TeamName == "radiant" && activity.Map.RadiantScore > activity.Map.DireScore {
				endGameText = "Won"
			}
			settings.Details = fmt.Sprintf("%s (%d - %d)", endGameText, activity.Map.RadiantScore, activity.Map.DireScore)
			settings.LargeImage = ImageDotaPostgame
		case string(DotaGameStatePreGame):
			settings.State = "Pre-Game"
			settings.Details = fixName(activity.Hero.Name)
			settings.LargeImage = getHeroImageUrl(activity.Hero.Name)
		case string(DotaGameStateTeamShowcase), string(DotaGameStateGameInProgress):
			settings.State = fmt.Sprintf("%d / %d / %d", activity.Player.Kills, activity.Player.Deaths, activity.Player.Assists)
			settings.Details = fmt.Sprintf("%s (LVL: %d)", fixName(activity.Hero.Name), activity.Hero.Level)
			settings.LargeImage = getHeroImageUrl(activity.Hero.Name)
			settings.LargeText = getItemsAsString(activity.Items)
			settings.SmallText = fmt.Sprintf("Match ID: %s\nPlaying for %s (%d - %d)", activity.Map.MatchID, activity.Player.TeamName, activity.Map.RadiantScore, activity.Map.DireScore)
			settings.Timestamps = &client.Timestamps{
				Start: &timestamp,
			}
		case string(DotaGameStateHeroSelection):
			settings.Details = "Hero Selection"
			settings.LargeImage = ImageDotaLoading
		case string(DotaGameStateStrategyTime):
			settings.Details = "Strategy discussion"
			settings.LargeImage = ImageDotaMap
		case string(DotaGameStateWaitForMapToLoad):
			settings.Details = "Loading on map..."
			settings.LargeImage = ImageDotaLoading
		case string(DotaGameStateWaitForPlayersToLoad):
			settings.Details = "Waiting for players..."
			settings.LargeImage = ImageDotaLoading
		default:
			settings.Details = "Unknown action:"
			settings.State = activity.Map.GameState
			settings.LargeImage = ImageDotaLogo
		}
	} else {
		settings.Details = "In Main Menu"
		settings.LargeImage = ImageDotaLogo
	}
	client.SetActivity(settings)
	lastUpdate = time.Now()
}
