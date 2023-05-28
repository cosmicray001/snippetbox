package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, _ = w.Write([]byte("Hello from snippetBox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Display a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		//w.WriteHeader(http.StatusMethodNotAllowed)
		//_, _ = w.Write([]byte("Method not allowed"))
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	_, _ = w.Write([]byte("Create a new snippet..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	log.Println("Starting server on :8000")
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
