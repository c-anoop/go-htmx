package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello, World!")
	// http.ServeFile(w, r, "index.html")
	data := struct {
		Message string
	}{
		Message: "Hello, World!",
	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

func handleTime(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Curent time from Go Server: " + time.Now().Format(time.RFC1123)))
}

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/time", handleTime)

	fmt.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
