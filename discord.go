package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/hugolgst/rich-go/client"
)

var loggedIn bool = false

func LogIn(clientId string) error {
	if loggedIn {
		return errors.New("ALREADY_LOGGED_IN")
	}
	err := client.Login(clientId)
	if err == nil {
		loggedIn = true
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

func Update() {
	timestamp := time.Unix(time.Now().Unix()-int64(activity.Map.ClockTime), 0)
	client.SetActivity(client.Activity{
		State:      fmt.Sprintf("(%d/%d/%d)", activity.Player.Kills, activity.Player.Deaths, activity.Player.Assists),
		Details:    fmt.Sprintf("%s (lvl %d)", fixName(activity.Hero.Name), activity.Hero.Level),
		LargeImage: getHeroImageUrl(activity.Hero.Name),
		LargeText:  getItemsAsString(),
		SmallImage: "https://cdn.discordapp.com/app-icons/356875988589740042/6b4b3fa4c83555d3008de69d33a60588",
		SmallText:  fmt.Sprintf("Radiant/Dire: %d/%d\nMatch ID: %s\nPlaying for %s", activity.Map.RadiantScore, activity.Map.DireScore, activity.Map.MatchID, activity.Player.TeamName),
		Timestamps: &client.Timestamps{
			Start: &timestamp,
		},
		// Timestamps: &client.Timestamps{
		// 	Start: &time.Unix(activity.Provider.Timestamp), //time.Unix(activity.Provider.Timestamp, 0),
		// },
	})
}
