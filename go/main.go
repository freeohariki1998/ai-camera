package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type DetectionEvent struct {
	Class      string  `json:"class"`      // 物体名
	Confidence float64 `json:"confidence"` // 確信度（0.0〜1.0）
	X          int     `json:"x"`          // バウンディングボックスの左上座標
	Y          int     `json:"y"`
	W          int     `json:"w"` // 幅・高さ
	H          int     `json:"h"`
	Timestamp  int64   `json:"timestamp"`
}

var repo *EventRepository
var hub *WebSocketHub

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "AI Camera Server is runnning")
}

func handleEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST only", http.StatusMethodNotAllowed)
		return
	}

	var event DetectionEvent
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if err := repo.Insert(event); err != nil {
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}

	jsonBytes, _ := json.Marshal(event)
	hub.broadcast <- jsonBytes

	log.Printf("Saved event: %+v\n", event)
	w.WriteHeader(http.StatusOK)
}
func main() {
	db := InitDB()
	repo = NewEventRepository(db)

	hub = NewWebSocketHub()
	go hub.Run()

	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/events", handleEvent)
	http.HandleFunc("/ws", hub.HandleWS)

	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
