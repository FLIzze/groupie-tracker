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

// type AllData struct {
//     People    []people
//     Locations Index
//     Dates     Index_date
// }

// type people struct {
//     Id           int      `json:"Id"`
//     Name         string   `json:"Name"`
//     Image        string   `json:"Image"`
//     Members      []string `json:"Members"`
//     FirstAlbum   string   `json:"FirstAlbum"`
//     CreationDate int      `json:"CreationDate"`
// }
// type Index struct {
//     Index []locations `json:"index"`
// }

// type Index_date struct {
//     Index []dates `json:"index"`
// }

// type dates struct {
//     Id   int      `json:"id"`
//     Date []string `json:"dates"`
// }

// type locations struct {
//     Id       int      `json:"id"`
//     Location []string `json:"locations"`
//     Date     string   `json:"dates"`
// }
// text := []people{}
// jsonErr := json.Unmarshal(body, &text)
// if jsonErr != nil {
// 	panic(jsonErr)
// }

// resp, err := http.Get(url2)
// if err != nil {
// 	print(err)
// }
// body2, readError := ioutil.ReadAll(resp.Body)
// if readError != nil {
// 	panic(readError)
// }

// txt := Index{}
// jsonError := json.Unmarshal(body2, &txt)
// if jsonError != nil {
// 	panic(jsonError)
// }

// res2, err := http.Get(url3)
// if err != nil {
// 	print(err)
// }
// body3, readErr := ioutil.ReadAll(res2.Body)
// if readErr != nil {
// 	panic(readErr)
// }

// date := Index_date{}
// jsonerr := json.Unmarshal(body3, &date)
// if jsonerr != nil {
// 	panic(jsonerr)
// }

// allData := AllData{
// 	People:    text,
// 	Locations: txt,
// 	Dates:     date,
// }

// tmpl := template.Must(template.ParseFiles("./static/index.html"))

// tmpl.Execute(w, allData)
// }
