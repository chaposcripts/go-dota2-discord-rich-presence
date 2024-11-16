package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type VersionInfo struct {
	Last      string              `json:"last_version"`
	Changelog map[string][]string `json:"changelog"`
}

func CheckForUpdates() {
	response, err := http.NewRequest("GET", "", nil)
	if err != nil {
		ShowToast("Error checking for updates", "Error in request: "+err.Error())
		return
	}
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		ShowToast("Error checking for updates", "Error in reading request: "+err.Error())
		return
	}
	var updateInfo VersionInfo
	err = json.Unmarshal(bytes, &updateInfo)
	if err != nil {
		ShowToast("Error checking for updates", "Error in decoding JSON: "+err.Error())
		return
	}
	if updateInfo.Last != version {
		ShowToast("New version available!", "Open github.com/chaposcripts/go-dota2-discord-rich-presence to download the latest version!")
	}
}
