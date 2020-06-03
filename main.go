package main

import (
	"go-module/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	myrouter := mux.NewRouter()

	// apis and handlers
	myrouter.HandleFunc("/checklogin", handlers.CheckLogin).Methods("POST")
	myrouter.HandleFunc("/checkuser", handlers.CheckCookie(handlers.EmptyHandler))

	staticDir := "/templates/"
	myrouter.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	http.ListenAndServe(":8000", myrouter)
}
