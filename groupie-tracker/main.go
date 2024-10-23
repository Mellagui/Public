package main

import (
	"fmt"
	"log"
	"net/http"
)

type Data struct {
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

	//tmpl, _ := template.ParseFiles("template.html")

}
