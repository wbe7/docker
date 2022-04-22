package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	env := os.Getenv("ENVIROMENT")
	fmt.Fprintf(w, "Hi there from %s, I love %s!", env, r.URL.Path[1:])
	log.Printf("Request from %s on Path /%s", r.RemoteAddr, r.URL.Path[1:])
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set")
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
