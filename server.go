package main

import (
	"fmt"
	"log"
	"net/http"
)

const addr string = ":8080"

func main() {

	http.HandleFunc("/health", healthHandler)
	fs := http.FileServer(http.Dir("./content"))
	http.Handle("/", fs)
	log.Printf("http server starting at %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}

}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}
