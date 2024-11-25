package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server 2: %s\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Server 2 running on port 8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
