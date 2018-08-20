// Name of package (FOLDER) always goes at the top - Main is the 'main' one that everything gets piped thorugh
package main

// Import paths similar to how you would import into a React project. These can either be other packages
// in your repo(game) or point to Go libraries(net/http, time, websocket)
import (
	"net/http"
	"time"

	"github.com/SadieF/ARGIPY-go-server/game"
	"github.com/gorilla/websocket"
)

// Upgrader upgrades HTTP server connection to the Websocket protocol.
// Client sends HTTP request to upgrade connection to the Websocket.
// If all good, the Server sends HTTP response agreeing to upgrade.
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

// Message struct - structs must have capitalised fields in order for them to be public,
// However the name of the struct can be uncapitalised and will be private to its specific package.
type Message struct {
	Type string
	Data interface{}
}

// Func main is the entry point for a Go application. It can't have any arguments or return any values.
// In any other function when you are taking in arguments and returning values, you must state the types of each argument
// and the types of each return value (often one of these will be 'error' if you are returning the error)
func main() {
	// An http.Handler wrapper is a function that has one input argument and one output argument, both of type http.Handler.
	// This handles the route for 'echo'
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil)

		// Variables in Go can either be described using var or const and declaring the type, OR using := which is shorthand for this.
		// If you are calling a function that returns only an error, you will often see it written like this,
		// Whereas if you are returning a value that you need to use, it will be written as:
		// res, err := FunctionHere()

		// res being whatever you want to call your variable.
		err := conn.WriteJSON(Message{
			Type: "fullupdate",
			Data: gameState,
		})
		// Error handling is done whenever you return an error. If you don't handle the errors, Go 'Panics'
		// (literally, it's called a panic) and stops running.
		if err != nil {
			return
		}
		time.Sleep(500 * time.Millisecond)

	})
	//This handles the route for '/' and serves the websockets file which is what creates a new websocket
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	})

	http.ListenAndServe(":8104", nil)
}
