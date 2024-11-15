package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func HandleGSI(port string) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var activity DotaGsiRequest
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			ShowToast("Server Error: Error reading request body", err.Error())
			fmt.Println("Error reading body", err.Error())
			return
		}
		err = json.Unmarshal(bytes, &activity)
		if err != nil {
			ShowToast("Server Error: Error decoding JSON", err.Error())
			fmt.Println("ERROR", err.Error())
			return
		}
		// fmt.Println("GSI Received, JSON:", string(bytes))
		fmt.Println("[Server] Request received")
		if checkbox.Checked() {
			Update(activity)
		} else {
			fmt.Println("Skipping update (disabled)")
		}
	})
	go func() {
		err := http.ListenAndServe(":"+port, nil)
		if err != nil {
			ShowMessageBox("Error", fmt.Sprintf("Error starting HTTP server!\nPort: %s\nError: %s\n\nClick \"Ok\" to close the application", port, err.Error()), MessageBoxTypeOk)
			os.Exit(1)
			return
		}
	}()
	ShowToast("Started, waiting for Dota2", "You can close or disable application in system tray")
	return nil
}
