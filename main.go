package main

import (
	"net/http"
	"time"

	"github.com/SadieF/ARGIPY-go-server/game"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var gameState = game.GameState{
	Players: map[string]game.Player{
		"a1234z": game.Player{
			Id:   "a1234z",
			Name: "jeremy",
		},
	},
	GridHeight: 200,
	GridWidth:  200,
	You:        "a1234z",
	Treasures: []game.Pos{
		game.Pos{1, 2},
	},
}

// Message ...
type Message struct {
	Type string
	Data interface{}
}

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil)

		err := conn.WriteJSON(Message{
			Type: "fullupdate",
			Data: gameState,
		})
		if err != nil {
			return
		}
		time.Sleep(500 * time.Millisecond)

	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	})

	http.ListenAndServe(":8080", nil)
}
