package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, _ = w.Write([]byte("Hello from snippetBox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	snippetID, err := strconv.Atoi(id)
	if err != nil || snippetID < 1 {
		http.Error(w, "Invalid Snippet ID", http.StatusNotFound)
		return
	}
	msg := fmt.Sprintf("Display a specific snippet with ID %d...", snippetID)
	_, _ = w.Write([]byte(msg))
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
