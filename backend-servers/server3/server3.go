package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server 3: %s\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Server 3 running on port 8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}
