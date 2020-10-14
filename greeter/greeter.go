package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request from %s", r.RemoteAddr)
	fmt.Fprintf(w, "<h1>%s %s!</h1>\n", os.Getenv("GREETING"), r.RemoteAddr)
}

func main() {
	log.Print("Server listening on port 80...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
