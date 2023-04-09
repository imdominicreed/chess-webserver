package handlers

import (
	"log"
	"net/http"

	"server/broadcaster"

	"github.com/gorilla/websocket"
	"github.com/notnil/chess"
)

var upgrader = websocket.Upgrader{}

var game = chess.NewGame()
var HandlerMap = map[string]func(http.ResponseWriter, *http.Request){"/move": HandleMove, "/gameupdate": HandlerGameUpdate}
var broadcast = broadcaster.NewBroadcaster()

func HandleMove(w http.ResponseWriter, r *http.Request) {
	move := r.FormValue("move")
	if err := game.MoveStr(move); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		log.Println("Invalid Move Recieved")

		return
	}
	log.Println("Valid Move Recieved and Broadcasting Update")

	go broadcast.Broadcast(move)
	log.Println("Broadcasted")

	w.WriteHeader(200)
	w.Write([]byte("Success"))
}

func HandlerGameUpdate(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade failed: ", err)
		return
	}
	log.Println("New Watcher, Adding Channel")

	id := broadcast.AddChannel()
	log.Println("New Channel Created")
	defer broadcast.DeleteChannel(id)
	defer conn.Close()

	for {
		log.Println("Awaiting Message")
		conn.WriteMessage(1, []byte(broadcast.AwaitMessage(id)))
		log.Println("Message Recieved")

	}
}
