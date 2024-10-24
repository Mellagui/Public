package main

import (
	"encoding/json" //////////
	"html/template"
	"log"
	"net/http"
)

type Data struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Felations    string   `json:"relations"`
}

func main() {

	http.HandleFunc("/home", handler)

	log.Println("Server start in : http://localhost:3000/home")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("Error:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	// api url
	url := "https://groupietrackers.herokuapp.com/api/artists"

	// http get request
	getResp, errG := http.Get(url)
	if errG != nil {
		log.Fatal("Error: http get request")
	}
	defer getResp.Body.Close()

	// check status is OK
	if getResp.StatusCode != 200 {
		log.Fatal("Error: statu code is not 200", getResp.StatusCode)
	}

	// decode the JSON response into a stract
	var apiRes []Data
	errj := json.NewDecoder(getResp.Body).Decode(&apiRes)
	if errj != nil {
		log.Fatalf("Error: json %v", errj)
	}
	//fmt.Println(apiRes)

	tmpl, _ := template.ParseFiles("template.html")
	data1 := apiRes
	//fmt.Println(data1)

	tmpl.Execute(w, data1)
}
