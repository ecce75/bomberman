package api

import (
	"bomberman/ws"
	"github.com/gorilla/mux"
	"net/http"
)

func Router(mux *mux.Router) {
	fs := http.FileServer(http.Dir("../frontend/public"))
	mux.PathPrefix("/").Handler(fs)

	mux.HandleFunc("/ws", ws.HandleConnections)
}
