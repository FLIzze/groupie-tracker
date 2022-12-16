package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type people struct {
	Id   int    `json:"Id"`
	Name string `json:"Name"`
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

	id1 := []people{}
	jsonErr := json.Unmarshal(body, &id1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(id1)
}

//requestURL := "https://groupietrackers.herokuapp.com/api/artists"
//request := := http.NewRequest()
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
