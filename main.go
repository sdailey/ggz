package main


import (
	"log"
	"net/http"
	
	// "route_handlers"
)

func main() {
	
	// initDb() ** 
	
	// loadGamesManifest()
	
	// startFileWatcher()
	
	// http.HandleFunc("/games/%a/%b", TakeToGame) **
	// http.HandleFunc("/save/%a/%b", SaveGameState) ** - JSON Web Token (JWT)
	
		// ->
			// /games/%a
				// >
	// http.HandleFunc("/library/%a/", TakeToGamePage) **
	
	// http.HandleFunc("/settings", TakeToSettings)
		// default: games local or remote
	
	// http.HandleFunc("/welcome", Welcome) // steps to import: 
	
	
	// start the server on port 8081
	log.Fatal(http.ListenAndServe(":8081", nil))
}