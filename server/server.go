package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type people struct {
	Id           int      `json:"Id"`
	Name         string   `json:"Name"`
	Image        string   `json:"Image"`
	Members      []string `json:"Members"`
	Location     string   `json:"Location"`
	FirstAlbum   string   `json:"FirstAlbum"`
	CreationDate int      `json:"CreationDate"`
}

var api people

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
	res, err := http.Get(url)
	if err != nil {
		print(err)
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	text := []people{}
	jsonErr := json.Unmarshal(body, &text)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	tmpl := template.Must(template.ParseFiles("./static/index.html"))
	fmt.Println(api.Id, api.Name)

	principalinput := r.FormValue("inputfilter")
	fmt.Println(principalinput)

	inputconcert := r.FormValue("concertsinput")
	fmt.Println(inputconcert)

	inputslide1 := r.FormValue("firstalbum")
	fmt.Println(inputslide1)

	inputslide2 := r.FormValue("created")
	fmt.Println(inputslide2)

	checkbox := r.FormValue("submit")
	fmt.Println(checkbox)

	tmpl.Execute(w, text)
}
