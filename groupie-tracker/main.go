package main

import (
	"encoding/json" //////////
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Artists struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relations struct {
	ID             int      `json:"id"`
	DatesLocations []string `json:"dateslocations"`
}

func main() {

	http.HandleFunc("/Home", handler)
	http.HandleFunc("/Artists", handlerCard)

	log.Println("Server start in : http://localhost:3000/Home")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("Error:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	// api url
	artistsURL := "https://groupietrackers.herokuapp.com/api/artists"

	// http get request
	getResp, errG := http.Get(artistsURL)
	if errG != nil {
		log.Fatal("Error: http get request ", errG)
	}
	defer getResp.Body.Close()

	// check status is OK
	if getResp.StatusCode != 200 {
		log.Fatal("Error: statu code is not 200", getResp.StatusCode)
	}

	// decode the JSON response into a stract
	var apiRes []Artists
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

func handlerCard(w http.ResponseWriter, r *http.Request) {

	// api url
	artistsURL := "https://groupietrackers.herokuapp.com/api/artists"

	// http get request
	getResp, errG := http.Get(artistsURL)
	if errG != nil {
		log.Fatal("Error: http get request")
	}
	defer getResp.Body.Close()

	// check status is OK
	if getResp.StatusCode != 200 {
		log.Fatal("Error: statu code is not 200", getResp.StatusCode) /////////////::
	}

	// decode the JSON response into a stract
	var apiRes []Artists
	errj := json.NewDecoder(getResp.Body).Decode(&apiRes) ///////////////:
	if errj != nil {
		log.Fatalf("Error: json %v", errj)
	}
	//fmt.Println(apiRes)

	tmpl, _ := template.ParseFiles("templateCard.html")
	data1 := apiRes[0]
	fmt.Println(data1)

	tmpl.Execute(w, data1)
}
