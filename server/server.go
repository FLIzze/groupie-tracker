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
	Id       int      `json:"Id"`
	Name     string   `json:"Name"`
	Status   string   `json:"Status"`
	Image    string   `json:"Image"`
	Members  []string `json:"Members"`
	Location string   `json:"Location"`
	First    string   `json:"First release"`
	Creation string   `json:"Creation"`
}

func main() {
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

	fmt.Println(text)
	http.HandleFunc("/", Handler_index)
	// //url of our funcs

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Print("Le Serveur dÃ©marre sur le port 8080\n")
	http.ListenAndServe(":8080", nil)
	//listening on port 8080

	type people struct {
		Id           int      `json:"Id"`
		Name         string   `json:"Name"`
		Image        string   `json:"Image"`
		Members      []string `json:"Members"`
		Location     string   `json:"Location"`
		FirstAlbum   string   `json:"FirstAlbum"`
		CreationDate string   `json:"CreationDate"`
	}
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

	fmt.Println(text)
}

func Handler_index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/index.html"))
	tmpl.Execute(w, r)
}
