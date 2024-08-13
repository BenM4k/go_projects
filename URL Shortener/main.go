package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var urlStore = make(map[string]string)

func shorten(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := r.Form.Get("url")
	shortURL := fmt.Sprintf("%x", md5.Sum([]byte(url)))[:5]
	urlStore[shortURL] = url
	fmt.Fprintf(w, "http://localhost:8080/%s\n", shortURL)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	originalUrl, ok := urlStore[vars["shortURL"]]
	if ok {
		http.Redirect(w, r, originalUrl, http.StatusMovedPermanently)
	} else {
		http.NotFound(w, r)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/create", shorten).Methods("POST")
	r.HandleFunc("/{shortURL}", redirect).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
