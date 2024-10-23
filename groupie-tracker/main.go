package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Data struct {
	Tit   string
	Hepe  string
	Forme string
	Bode  string
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
	fmt.Fprintln(w, "Hello World !!")

	tmpl, _ := template.ParseFiles("template.html")
	data := Data{
		Tit:   "title",
		Hepe:  "heading",
		Forme: "formulaire",
		Bode:  "bodying",
	}
	tmpl.Execute(w, data)
}
