package handlers

import (
	"net/http"

	"github.com/notnil/chess"
)

var game = chess.NewGame()

func HandleMove(w http.ResponseWriter, r *http.Request) {
	move := r.FormValue("move")
	if err := game.MoveStr(move); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("Success"))
}
