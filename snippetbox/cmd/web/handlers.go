package main 

import (
	"fmt"
	"log"
	"html/template"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	//404 invalid paths
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ts, err := template.ParseFiles("./ui/html/home.page.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

	//w.Write([]byte("Hello from Snippetbox"))
}

//add showSnippet handler function 
func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	// Use the fmt.Fprintf() function to interpolate the id value with our response 
	// and write it to the http.ResponseWriter.
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

//add a createSnippet Handler
func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Write([]byte("Create a new snippet..."))
} 