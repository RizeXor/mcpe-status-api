package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ZeroErrors/go-bedrockping"
	"github.com/gorilla/mux"
)

// ServerStatus : Status of server
type ServerStatus struct {
	Online      bool `json:"online"`
	PlayerCount int  `json:"player_count"`
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var s = &ServerStatus{false, 0}

	ip := "34.93.14.207"
	resp, err := bedrockping.Query(ip+":"+"19132", 500*time.Millisecond, 150*time.Millisecond)

	if err != nil {
		json.NewEncoder(w).Encode(s)
		return
	}

	s = &ServerStatus{true, resp.PlayerCount}
	fmt.Printf("%d/%d players are online.\n", resp.PlayerCount, resp.MaxPlayers)

	json.NewEncoder(w).Encode(s)
}

func main() {
	// Init router
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/status", getStatus)
	http.ListenAndServe(":8000", r)
}
