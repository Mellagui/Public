package main

import (
	"encoding/json" //////////
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	LocationsLink    string   `json:"locations"`
	ConcertDatesLink string   `json:"concertDates"`
	RelationsLink    string   `json:"relations"`
	Locations []string
	Dates []string
	Relations map[string][]string
}

var artists []Artist

func init() {
	//Curling Data
	fmt.Println("Curling data...")
	getArtists()
	getSubData()
	fmt.Println("data obtained successfully")
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

	tmpl, _ := template.ParseFiles("template.html")
	data1 := artists
	//fmt.Println(data1)

	tmpl.Execute(w, data1)
}

func handlerCard(w http.ResponseWriter, r *http.Request) {

	tmpl, _ := template.ParseFiles("templateCard.html")

	if err := r.ParseForm(); err != nil {
		http.Error(w, "500 Internal sever error - error parsing form", 500)
		fmt.Println(err)
	}
	idString := r.Form.Get("id")
	id, err := strconv.Atoi(idString)

	if err != nil || id >= len(artists) {
		http.Error(w, "404 - Not Found", 404)
		fmt.Println("error getting id")
		return
	}

	data1 := artists[id-1]
	tmpl.Execute(w, data1)
}

func getArtists() {
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
	errj := json.NewDecoder(getResp.Body).Decode(&artists)
	if errj != nil {
		log.Fatalf("Error: json %v", errj)
	}
}

func getSubData() {
	urls := []string{
	 "https://groupietrackers.herokuapp.com/api/locations",
	 "https://groupietrackers.herokuapp.com/api/dates", 
	 "https://groupietrackers.herokuapp.com/api/relation",
	}
	result := make([]map[string][]map[string]any, 3)
	for i := range urls {
		// http get request
		getResp, errG := http.Get(urls[i])
		if errG != nil {
			log.Fatal("Error: http get request")
		}
		defer getResp.Body.Close()

		// check status is OK
		if getResp.StatusCode != 200 {
			log.Fatal("Error: status code is not 200", getResp.StatusCode) /////////////::
		}

		// decode the JSON response into a stract
		errj := json.NewDecoder(getResp.Body).Decode(&result[i]) ///////////////:
		// if i == 1 {
		// 	errj = json.NewDecoder(getResp.Body).Decode(&dates) 
		// }
		if errj != nil {
			log.Fatalf("Error: json %v", errj)
		}
	}
	
	for i := range artists {
		// Assigning dates :
		artists[i].Locations = interfaceToStringSlice(result[0]["index"][i]["locations"])

		// Assigning dates :
		artists[i].Dates = interfaceToStringSlice(result[1]["index"][i]["dates"])
		
		// Assigning relations :
		artists[i].Relations = interfaceToMap(result[2]["index"][i]["datesLocations"])
	}
}


func interfaceToMap(input any) (map[string][]string) {
    // First, try to assert input as map[string]interface{}
    interfaceMap := input.(map[string]any)

    // Create a new map[string]string to hold the converted values
    stringMap := make(map[string][]string)

    // Loop through each element and try to convert it to a string
    for key, value := range interfaceMap {

		slice := value.([]any)

		dates := make([]string, len(slice))

		for i, v := range slice {
			str := v.(string)
			dates[i] = str
		}

        stringMap[key] = dates
    }

    return stringMap
}

func interfaceToStringSlice(input any) ([]string) {
    // First, try to assert input as []any
    interfaceSlice, ok := input.([]any)
    if !ok {
		fmt.Println("input is not a []interface{}")
        return nil
    }

    // Create a new []string slice to hold the converted values
    stringSlice := make([]string, len(interfaceSlice))

    // Loop through each element and try to convert it to a string
    for i, v := range interfaceSlice {
        str := v.(string)
        stringSlice[i] = str
    }

    return stringSlice
}
