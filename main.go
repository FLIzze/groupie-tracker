package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type people struct {
	id   int    `json:"Id"`
	name string `json:"Name"`
}

func main() {
	url := "https://groupietrackers.herokuapp.com/api/artists"

	spaceClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	id1 := people{}
	jsonErr := json.Unmarshal(body, &id1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(id1.name)
}

//requestURL := "https://groupietrackers.herokuapp.com/api/artists"
//request := new XMLHttpRequest()
//contenu := requests.get(requestURL)
//list := contenu.json()
//request.send()

//textBytes := []byte(text)

//people1 := people{}
//err := json.Unmarshal(textBytes, &people1)
//if err != nil {
//fmt.Println(err)
//return
//}
//fmt.Println(people1.Name)
//}
