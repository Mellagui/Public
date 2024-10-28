package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	web "asciiArtWeb/asciiArtFs" // This imports the custom ASCII art generation package from your project.
	// Helps with HTML templates that are used to render the webpage.
	// Used for logging errors and messages to the console.
	// Allows the creation of an HTTP server.
)

type Data struct {
	Text     string        // Text: Stores the input text.
	Banner   string        // Banner: Stores the font style for ASCII art.
	AsciiArt template.HTML // AsciiArt: Stores the generated ASCII art.
}

func main() {
	// Register the handler for the root URL
	http.HandleFunc("/", AppHandler)

	// Start the web server
	log.Println("Starting server on http://localhost:8080/")
	err := http.ListenAndServe(":8080", nil) // nil: use default miltiplixer / use nil when use (http.HandleFunc("/", AppHandler))
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}
}

func AsciiArtMaker(text string, banner string) (string, []error) {
	errs := []error{}
	if banner == "all" {
		AsciiArt1, err := web.AsciiArtFs(text, "standard")
		errs = append(errs, err)
		AsciiArt2, err := web.AsciiArtFs(text, "shadow")
		errs = append(errs, err)
		AsciiArt3, err := web.AsciiArtFs(text, "thinkertoy")
		errs = append(errs, err)
		return AsciiArt1 + AsciiArt2 + AsciiArt3, errs
	}
	AsciiArt, err := web.AsciiArtFs(text, banner)
	return AsciiArt, []error{err}
}

func AppHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		Get(w, r)
	} else if r.Method == "POST" {
		Post(w, r)
	} else {
		http.Error(w, "405 - Method Not Allowed", 405)
	}
}

func Get(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" { /////...
		http.Error(w, "404 - Not Found", 404)
		return
	}
	// http.ServeFile(w, r, "template.html")
	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		http.Error(w, "404 - Not Found", 404)
		return
	}
	data := Data{}
	tmpl.Execute(w, data)
}

func Post(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		http.Error(w, "404- Not Found", 404)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "500 - Internal Server Error", 500)
		log.Fatalln(err)
	}

	fmt.Println(r.Form)

	// 1
	// text := r.Form.Get("text")
	// banner := r.Form.Get("banner")

	// 2
	text := r.Form["text"][0]
	banner := r.Form["banner"][0]

	if len(text) == 0 || len(banner) == 0 {
		http.Error(w, "400 - Bad Request", 400)
		return
	}

	asciiArt, errs := AsciiArtMaker(text, banner)
	tmpl, err := template.ParseFiles("template.html")
	errs = append(errs, err)

	// Handing template err and AsciiConverter errs
	for i := range errs {
		if errs[i] != nil {
			notFound := fmt.Errorf("NotFound")
			if errs[i] == notFound {
				http.Error(w, "404 - Not Found", 404)
			} else {
				http.Error(w, "500 - Internal Server Error", 500)
			}
			return
		}
	}

	data := Data{
		Text:     text,
		Banner:   banner,
		AsciiArt: template.HTML(asciiArt),
	}
	tmpl.Execute(w, data)
}
