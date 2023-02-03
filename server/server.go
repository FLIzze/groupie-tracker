package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type people struct {
	Id           int      `json:"Id"`
	Name         string   `json:"Name"`
	Image        string   `json:"Image"`
	Members      []string `json:"Members"`
	FirstAlbum   string   `json:"FirstAlbum"`
	CreationDate int      `json:"CreationDate"`
}

func main() {
	http.HandleFunc("/", Handler_main)
	// //url of our funcs
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Print("Le Serveur dÃ©marre sur le port 8080\n")
	http.ListenAndServe(":8080", nil)
	//listening on port 8080
}
func Handler_main(w http.ResponseWriter, r *http.Request) {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	//var api people
	res, err := http.Get(url)
	if err != nil {
		print(err)
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		panic(readErr)
	}
	text := []people{}
	jsonErr := json.Unmarshal(body, &text)
	if jsonErr != nil {
		panic(jsonErr)
	}
	tmpl := template.Must(template.ParseFiles("./static/index.html"))
	tmpl.Execute(w, text)
}
