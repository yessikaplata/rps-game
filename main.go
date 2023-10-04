package main

import (
	"log"
	"net/http"

	"github.com/yessikaplata/rps-game/handler"
)

func main() {
	// Create router
	port := ":8080"
	router := http.NewServeMux()

	// Statics files (css)
	fs := http.FileServer(http.Dir("static"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	router.HandleFunc("/", handler.Index)
	router.HandleFunc("/about", handler.About)
	router.HandleFunc("/new", handler.NewGame)
	router.HandleFunc("/game", handler.Game)
	router.HandleFunc("/play", handler.Play)

	// Start server
	log.Printf("Server listening in http://localhost%s\n", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal(err)
	}

}
