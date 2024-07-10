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
	w.Write([]byte("home"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("snippetView"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(405)
		w.Write([]byte("method not allowed"))
		return
	}

	w.Write([]byte("snippetCreate"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Listening on port 7000")

	err := http.ListenAndServe(":7000", mux)
	log.Fatal(err)
}
