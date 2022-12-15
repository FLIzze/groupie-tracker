package main

import (
	"fmt"
	"net/http"
)

func main() {
	// http.HandleFunc("/", Handler_index)
	// http.HandleFunc("/login", Handler_login)
	// http.HandleFunc("/win", Handler_win)
	// http.HandleFunc("/loose", Handler_loose)
	// //url of our funcs

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Print("Le Serveur dÃ©marre sur le port 8080\n")
	http.ListenAndServe(":8080", nil)
	//listening on port 8080
}
