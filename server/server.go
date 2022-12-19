package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	goApi()

	http.HandleFunc("/", Handler_index)
	// //url of our funcs

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Print("Le Serveur dÃ©marre sur le port 8080\n")
	http.ListenAndServe(":8080", nil)
	//listening on port 8080
}

func Handler_index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/index.html"))
	tmpl.Execute(w, r)
}

func goApi() {
	type people struct {
		Id           int      `json:"Id"`
		Name         string   `json:"Name"`
		Image        string   `json:"Image"`
		Members      []string `json:"Members"`
		Location     string   `json:"Location"`
		FirstAlbum   string   `json:"FirstAlbum"`
		CreationDate int      `json:"CreationDate"`
	}

	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
}
