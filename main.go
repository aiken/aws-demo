package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleMainPage)

	log.Println("Starting webserver on :8082")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatal("http.ListendAndServer() failed with %s\n", err)
	}
}

func handleMainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Hello World, Aiken\n")
}
