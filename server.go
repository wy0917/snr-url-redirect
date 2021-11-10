package main

import (
	"log"
	"net/http"
)

var RedirectURL string

// Accept a request and silently redirect to customized url
func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, RedirectURL, 301)
}

func main() {
	http.HandleFunc("/", redirect)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
