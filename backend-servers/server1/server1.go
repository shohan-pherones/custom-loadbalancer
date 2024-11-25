package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Server 1: %s\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Server 1 running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
