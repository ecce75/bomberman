package api

import (
	"bomberman/ws"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func Router(mux *mux.Router) {
	mux.PathPrefix("/").Handler(fileHandler())
	// mux.PathPrefix("/frontend/src").Handler(serveStaticFiles())

	mux.HandleFunc("/ws", ws.HandleConnections)
}

// func serveStaticFiles() http.Handler {
//
// }

func fileHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the requested file is a JavaScript file
		if strings.HasSuffix(r.URL.Path, ".js") {
			// Set Content-Type header for JavaScript files
			w.Header().Set("Content-Type", "application/javascript")
		}
		// Serve files using http.FileServer, it correctly serves other files
		http.FileServer(http.Dir("../frontend/src")).ServeHTTP(w, r)
	})
}
