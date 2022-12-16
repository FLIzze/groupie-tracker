package main

import (
	"encoding/json"
	"fmt"
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
	CreationDate string   `json:"CreationDate"`
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
}
