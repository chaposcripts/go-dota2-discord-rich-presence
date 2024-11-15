package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func HandleGSI(port string) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var activity DotaGsiRequest
		fmt.Println("GSI Received")
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error reading body", err.Error())
			return
		}
		err = json.Unmarshal(bytes, &activity)
		if err != nil {
			fmt.Println("ERROR", err.Error())
			return
		}
		// fmt.Println(fmt.Sprintf("GameState: %s, player activity: %s", activity.Map.GameState, activity.Player.Activity))
		fmt.Println(string(bytes))
		if checkbox.Checked() {
			Update(activity)
		} else {
			fmt.Println("Skipping update (disabled)")
		}
	})

	func() {
		err := http.ListenAndServe(":"+port, nil)
		if err != nil {
			fmt.Println("Error starting server:", err.Error())
			return
		}
	}()
	return nil
}
